# Copyright 2021 Google LLC. All Rights Reserved.
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
from google3.cloud.graphite.mmv2.services.google.monitoring import alert_policy_pb2
from google3.cloud.graphite.mmv2.services.google.monitoring import alert_policy_pb2_grpc

from typing import List


class AlertPolicy(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        documentation: dict = None,
        user_labels: dict = None,
        conditions: list = None,
        combiner: str = None,
        disabled: bool = None,
        enabled: dict = None,
        validity: dict = None,
        notification_channels: list = None,
        creation_record: dict = None,
        mutation_record: dict = None,
        incident_strategy: dict = None,
        metadata: dict = None,
        project: str = None,
        service_account_file: str = "",
    ):

        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.documentation = documentation
        self.user_labels = user_labels
        self.conditions = conditions
        self.combiner = combiner
        self.disabled = disabled
        self.enabled = enabled
        self.validity = validity
        self.notification_channels = notification_channels
        self.creation_record = creation_record
        self.mutation_record = mutation_record
        self.incident_strategy = incident_strategy
        self.metadata = metadata
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = alert_policy_pb2_grpc.MonitoringAlertPolicyServiceStub(channel.Channel())
        request = alert_policy_pb2.ApplyMonitoringAlertPolicyRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if AlertPolicyDocumentation.to_proto(self.documentation):
            request.resource.documentation.CopyFrom(
                AlertPolicyDocumentation.to_proto(self.documentation)
            )
        else:
            request.resource.ClearField("documentation")
        if Primitive.to_proto(self.user_labels):
            request.resource.user_labels = Primitive.to_proto(self.user_labels)

        if AlertPolicyConditionsArray.to_proto(self.conditions):
            request.resource.conditions.extend(
                AlertPolicyConditionsArray.to_proto(self.conditions)
            )
        if AlertPolicyCombinerEnum.to_proto(self.combiner):
            request.resource.combiner = AlertPolicyCombinerEnum.to_proto(self.combiner)

        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if AlertPolicyEnabled.to_proto(self.enabled):
            request.resource.enabled.CopyFrom(AlertPolicyEnabled.to_proto(self.enabled))
        else:
            request.resource.ClearField("enabled")
        if AlertPolicyValidity.to_proto(self.validity):
            request.resource.validity.CopyFrom(
                AlertPolicyValidity.to_proto(self.validity)
            )
        else:
            request.resource.ClearField("validity")
        if Primitive.to_proto(self.notification_channels):
            request.resource.notification_channels.extend(
                Primitive.to_proto(self.notification_channels)
            )
        if AlertPolicyCreationRecord.to_proto(self.creation_record):
            request.resource.creation_record.CopyFrom(
                AlertPolicyCreationRecord.to_proto(self.creation_record)
            )
        else:
            request.resource.ClearField("creation_record")
        if AlertPolicyMutationRecord.to_proto(self.mutation_record):
            request.resource.mutation_record.CopyFrom(
                AlertPolicyMutationRecord.to_proto(self.mutation_record)
            )
        else:
            request.resource.ClearField("mutation_record")
        if AlertPolicyIncidentStrategy.to_proto(self.incident_strategy):
            request.resource.incident_strategy.CopyFrom(
                AlertPolicyIncidentStrategy.to_proto(self.incident_strategy)
            )
        else:
            request.resource.ClearField("incident_strategy")
        if AlertPolicyMetadata.to_proto(self.metadata):
            request.resource.metadata.CopyFrom(
                AlertPolicyMetadata.to_proto(self.metadata)
            )
        else:
            request.resource.ClearField("metadata")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyMonitoringAlertPolicy(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.documentation = AlertPolicyDocumentation.from_proto(response.documentation)
        self.user_labels = Primitive.from_proto(response.user_labels)
        self.conditions = AlertPolicyConditionsArray.from_proto(response.conditions)
        self.combiner = AlertPolicyCombinerEnum.from_proto(response.combiner)
        self.disabled = Primitive.from_proto(response.disabled)
        self.enabled = AlertPolicyEnabled.from_proto(response.enabled)
        self.validity = AlertPolicyValidity.from_proto(response.validity)
        self.notification_channels = Primitive.from_proto(
            response.notification_channels
        )
        self.creation_record = AlertPolicyCreationRecord.from_proto(
            response.creation_record
        )
        self.mutation_record = AlertPolicyMutationRecord.from_proto(
            response.mutation_record
        )
        self.incident_strategy = AlertPolicyIncidentStrategy.from_proto(
            response.incident_strategy
        )
        self.metadata = AlertPolicyMetadata.from_proto(response.metadata)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = alert_policy_pb2_grpc.MonitoringAlertPolicyServiceStub(channel.Channel())
        request = alert_policy_pb2.DeleteMonitoringAlertPolicyRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if AlertPolicyDocumentation.to_proto(self.documentation):
            request.resource.documentation.CopyFrom(
                AlertPolicyDocumentation.to_proto(self.documentation)
            )
        else:
            request.resource.ClearField("documentation")
        if Primitive.to_proto(self.user_labels):
            request.resource.user_labels = Primitive.to_proto(self.user_labels)

        if AlertPolicyConditionsArray.to_proto(self.conditions):
            request.resource.conditions.extend(
                AlertPolicyConditionsArray.to_proto(self.conditions)
            )
        if AlertPolicyCombinerEnum.to_proto(self.combiner):
            request.resource.combiner = AlertPolicyCombinerEnum.to_proto(self.combiner)

        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if AlertPolicyEnabled.to_proto(self.enabled):
            request.resource.enabled.CopyFrom(AlertPolicyEnabled.to_proto(self.enabled))
        else:
            request.resource.ClearField("enabled")
        if AlertPolicyValidity.to_proto(self.validity):
            request.resource.validity.CopyFrom(
                AlertPolicyValidity.to_proto(self.validity)
            )
        else:
            request.resource.ClearField("validity")
        if Primitive.to_proto(self.notification_channels):
            request.resource.notification_channels.extend(
                Primitive.to_proto(self.notification_channels)
            )
        if AlertPolicyCreationRecord.to_proto(self.creation_record):
            request.resource.creation_record.CopyFrom(
                AlertPolicyCreationRecord.to_proto(self.creation_record)
            )
        else:
            request.resource.ClearField("creation_record")
        if AlertPolicyMutationRecord.to_proto(self.mutation_record):
            request.resource.mutation_record.CopyFrom(
                AlertPolicyMutationRecord.to_proto(self.mutation_record)
            )
        else:
            request.resource.ClearField("mutation_record")
        if AlertPolicyIncidentStrategy.to_proto(self.incident_strategy):
            request.resource.incident_strategy.CopyFrom(
                AlertPolicyIncidentStrategy.to_proto(self.incident_strategy)
            )
        else:
            request.resource.ClearField("incident_strategy")
        if AlertPolicyMetadata.to_proto(self.metadata):
            request.resource.metadata.CopyFrom(
                AlertPolicyMetadata.to_proto(self.metadata)
            )
        else:
            request.resource.ClearField("metadata")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteMonitoringAlertPolicy(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = alert_policy_pb2_grpc.MonitoringAlertPolicyServiceStub(channel.Channel())
        request = alert_policy_pb2.ListMonitoringAlertPolicyRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListMonitoringAlertPolicy(request).items

    def to_proto(self):
        resource = alert_policy_pb2.MonitoringAlertPolicy()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if AlertPolicyDocumentation.to_proto(self.documentation):
            resource.documentation.CopyFrom(
                AlertPolicyDocumentation.to_proto(self.documentation)
            )
        else:
            resource.ClearField("documentation")
        if Primitive.to_proto(self.user_labels):
            resource.user_labels = Primitive.to_proto(self.user_labels)
        if AlertPolicyConditionsArray.to_proto(self.conditions):
            resource.conditions.extend(
                AlertPolicyConditionsArray.to_proto(self.conditions)
            )
        if AlertPolicyCombinerEnum.to_proto(self.combiner):
            resource.combiner = AlertPolicyCombinerEnum.to_proto(self.combiner)
        if Primitive.to_proto(self.disabled):
            resource.disabled = Primitive.to_proto(self.disabled)
        if AlertPolicyEnabled.to_proto(self.enabled):
            resource.enabled.CopyFrom(AlertPolicyEnabled.to_proto(self.enabled))
        else:
            resource.ClearField("enabled")
        if AlertPolicyValidity.to_proto(self.validity):
            resource.validity.CopyFrom(AlertPolicyValidity.to_proto(self.validity))
        else:
            resource.ClearField("validity")
        if Primitive.to_proto(self.notification_channels):
            resource.notification_channels.extend(
                Primitive.to_proto(self.notification_channels)
            )
        if AlertPolicyCreationRecord.to_proto(self.creation_record):
            resource.creation_record.CopyFrom(
                AlertPolicyCreationRecord.to_proto(self.creation_record)
            )
        else:
            resource.ClearField("creation_record")
        if AlertPolicyMutationRecord.to_proto(self.mutation_record):
            resource.mutation_record.CopyFrom(
                AlertPolicyMutationRecord.to_proto(self.mutation_record)
            )
        else:
            resource.ClearField("mutation_record")
        if AlertPolicyIncidentStrategy.to_proto(self.incident_strategy):
            resource.incident_strategy.CopyFrom(
                AlertPolicyIncidentStrategy.to_proto(self.incident_strategy)
            )
        else:
            resource.ClearField("incident_strategy")
        if AlertPolicyMetadata.to_proto(self.metadata):
            resource.metadata.CopyFrom(AlertPolicyMetadata.to_proto(self.metadata))
        else:
            resource.ClearField("metadata")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class AlertPolicyDocumentation(object):
    def __init__(self, content: str = None, mime_type: str = None):
        self.content = content
        self.mime_type = mime_type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyDocumentation()
        if Primitive.to_proto(resource.content):
            res.content = Primitive.to_proto(resource.content)
        if Primitive.to_proto(resource.mime_type):
            res.mime_type = Primitive.to_proto(resource.mime_type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyDocumentation(
            content=Primitive.from_proto(resource.content),
            mime_type=Primitive.from_proto(resource.mime_type),
        )


class AlertPolicyDocumentationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AlertPolicyDocumentation.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AlertPolicyDocumentation.from_proto(i) for i in resources]


class AlertPolicyConditions(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        resource_state_filter: str = None,
        condition_threshold: dict = None,
        condition_absent: dict = None,
        condition_matched_log: dict = None,
        condition_cluster_outlier: dict = None,
        condition_rate: dict = None,
        condition_up_mon: dict = None,
        condition_process_count: dict = None,
        condition_time_series_query_language: dict = None,
        condition_monitoring_query_language: dict = None,
    ):
        self.name = name
        self.display_name = display_name
        self.resource_state_filter = resource_state_filter
        self.condition_threshold = condition_threshold
        self.condition_absent = condition_absent
        self.condition_matched_log = condition_matched_log
        self.condition_cluster_outlier = condition_cluster_outlier
        self.condition_rate = condition_rate
        self.condition_up_mon = condition_up_mon
        self.condition_process_count = condition_process_count
        self.condition_time_series_query_language = condition_time_series_query_language
        self.condition_monitoring_query_language = condition_monitoring_query_language

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyConditions()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.display_name):
            res.display_name = Primitive.to_proto(resource.display_name)
        if AlertPolicyConditionsResourceStateFilterEnum.to_proto(
            resource.resource_state_filter
        ):
            res.resource_state_filter = AlertPolicyConditionsResourceStateFilterEnum.to_proto(
                resource.resource_state_filter
            )
        if AlertPolicyConditionsConditionThreshold.to_proto(
            resource.condition_threshold
        ):
            res.condition_threshold.CopyFrom(
                AlertPolicyConditionsConditionThreshold.to_proto(
                    resource.condition_threshold
                )
            )
        else:
            res.ClearField("condition_threshold")
        if AlertPolicyConditionsConditionAbsent.to_proto(resource.condition_absent):
            res.condition_absent.CopyFrom(
                AlertPolicyConditionsConditionAbsent.to_proto(resource.condition_absent)
            )
        else:
            res.ClearField("condition_absent")
        if AlertPolicyConditionsConditionMatchedLog.to_proto(
            resource.condition_matched_log
        ):
            res.condition_matched_log.CopyFrom(
                AlertPolicyConditionsConditionMatchedLog.to_proto(
                    resource.condition_matched_log
                )
            )
        else:
            res.ClearField("condition_matched_log")
        if AlertPolicyConditionsConditionClusterOutlier.to_proto(
            resource.condition_cluster_outlier
        ):
            res.condition_cluster_outlier.CopyFrom(
                AlertPolicyConditionsConditionClusterOutlier.to_proto(
                    resource.condition_cluster_outlier
                )
            )
        else:
            res.ClearField("condition_cluster_outlier")
        if AlertPolicyConditionsConditionRate.to_proto(resource.condition_rate):
            res.condition_rate.CopyFrom(
                AlertPolicyConditionsConditionRate.to_proto(resource.condition_rate)
            )
        else:
            res.ClearField("condition_rate")
        if AlertPolicyConditionsConditionUpMon.to_proto(resource.condition_up_mon):
            res.condition_up_mon.CopyFrom(
                AlertPolicyConditionsConditionUpMon.to_proto(resource.condition_up_mon)
            )
        else:
            res.ClearField("condition_up_mon")
        if AlertPolicyConditionsConditionProcessCount.to_proto(
            resource.condition_process_count
        ):
            res.condition_process_count.CopyFrom(
                AlertPolicyConditionsConditionProcessCount.to_proto(
                    resource.condition_process_count
                )
            )
        else:
            res.ClearField("condition_process_count")
        if AlertPolicyConditionsConditionTimeSeriesQueryLanguage.to_proto(
            resource.condition_time_series_query_language
        ):
            res.condition_time_series_query_language.CopyFrom(
                AlertPolicyConditionsConditionTimeSeriesQueryLanguage.to_proto(
                    resource.condition_time_series_query_language
                )
            )
        else:
            res.ClearField("condition_time_series_query_language")
        if AlertPolicyConditionsConditionMonitoringQueryLanguage.to_proto(
            resource.condition_monitoring_query_language
        ):
            res.condition_monitoring_query_language.CopyFrom(
                AlertPolicyConditionsConditionMonitoringQueryLanguage.to_proto(
                    resource.condition_monitoring_query_language
                )
            )
        else:
            res.ClearField("condition_monitoring_query_language")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditions(
            name=Primitive.from_proto(resource.name),
            display_name=Primitive.from_proto(resource.display_name),
            resource_state_filter=AlertPolicyConditionsResourceStateFilterEnum.from_proto(
                resource.resource_state_filter
            ),
            condition_threshold=AlertPolicyConditionsConditionThreshold.from_proto(
                resource.condition_threshold
            ),
            condition_absent=AlertPolicyConditionsConditionAbsent.from_proto(
                resource.condition_absent
            ),
            condition_matched_log=AlertPolicyConditionsConditionMatchedLog.from_proto(
                resource.condition_matched_log
            ),
            condition_cluster_outlier=AlertPolicyConditionsConditionClusterOutlier.from_proto(
                resource.condition_cluster_outlier
            ),
            condition_rate=AlertPolicyConditionsConditionRate.from_proto(
                resource.condition_rate
            ),
            condition_up_mon=AlertPolicyConditionsConditionUpMon.from_proto(
                resource.condition_up_mon
            ),
            condition_process_count=AlertPolicyConditionsConditionProcessCount.from_proto(
                resource.condition_process_count
            ),
            condition_time_series_query_language=AlertPolicyConditionsConditionTimeSeriesQueryLanguage.from_proto(
                resource.condition_time_series_query_language
            ),
            condition_monitoring_query_language=AlertPolicyConditionsConditionMonitoringQueryLanguage.from_proto(
                resource.condition_monitoring_query_language
            ),
        )


class AlertPolicyConditionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AlertPolicyConditions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AlertPolicyConditions.from_proto(i) for i in resources]


class AlertPolicyConditionsConditionThreshold(object):
    def __init__(
        self,
        filter: str = None,
        aggregations: list = None,
        denominator_filter: str = None,
        denominator_aggregations: list = None,
        comparison: str = None,
        threshold_value: float = None,
        duration: str = None,
        trigger: dict = None,
    ):
        self.filter = filter
        self.aggregations = aggregations
        self.denominator_filter = denominator_filter
        self.denominator_aggregations = denominator_aggregations
        self.comparison = comparison
        self.threshold_value = threshold_value
        self.duration = duration
        self.trigger = trigger

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThreshold()
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if AlertPolicyConditionsConditionThresholdAggregationsArray.to_proto(
            resource.aggregations
        ):
            res.aggregations.extend(
                AlertPolicyConditionsConditionThresholdAggregationsArray.to_proto(
                    resource.aggregations
                )
            )
        if Primitive.to_proto(resource.denominator_filter):
            res.denominator_filter = Primitive.to_proto(resource.denominator_filter)
        if AlertPolicyConditionsConditionThresholdDenominatorAggregationsArray.to_proto(
            resource.denominator_aggregations
        ):
            res.denominator_aggregations.extend(
                AlertPolicyConditionsConditionThresholdDenominatorAggregationsArray.to_proto(
                    resource.denominator_aggregations
                )
            )
        if AlertPolicyConditionsConditionThresholdComparisonEnum.to_proto(
            resource.comparison
        ):
            res.comparison = AlertPolicyConditionsConditionThresholdComparisonEnum.to_proto(
                resource.comparison
            )
        if Primitive.to_proto(resource.threshold_value):
            res.threshold_value = Primitive.to_proto(resource.threshold_value)
        if Primitive.to_proto(resource.duration):
            res.duration = Primitive.to_proto(resource.duration)
        if AlertPolicyConditionsConditionThresholdTrigger.to_proto(resource.trigger):
            res.trigger.CopyFrom(
                AlertPolicyConditionsConditionThresholdTrigger.to_proto(
                    resource.trigger
                )
            )
        else:
            res.ClearField("trigger")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionThreshold(
            filter=Primitive.from_proto(resource.filter),
            aggregations=AlertPolicyConditionsConditionThresholdAggregationsArray.from_proto(
                resource.aggregations
            ),
            denominator_filter=Primitive.from_proto(resource.denominator_filter),
            denominator_aggregations=AlertPolicyConditionsConditionThresholdDenominatorAggregationsArray.from_proto(
                resource.denominator_aggregations
            ),
            comparison=AlertPolicyConditionsConditionThresholdComparisonEnum.from_proto(
                resource.comparison
            ),
            threshold_value=Primitive.from_proto(resource.threshold_value),
            duration=Primitive.from_proto(resource.duration),
            trigger=AlertPolicyConditionsConditionThresholdTrigger.from_proto(
                resource.trigger
            ),
        )


class AlertPolicyConditionsConditionThresholdArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AlertPolicyConditionsConditionThreshold.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionThreshold.from_proto(i) for i in resources
        ]


class AlertPolicyConditionsConditionThresholdAggregations(object):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
        reduce_fraction_less_than_params: dict = None,
        reduce_make_distribution_params: dict = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields
        self.reduce_fraction_less_than_params = reduce_fraction_less_than_params
        self.reduce_make_distribution_params = reduce_make_distribution_params

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdAggregations()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        if AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams.to_proto(
            resource.reduce_fraction_less_than_params
        ):
            res.reduce_fraction_less_than_params.CopyFrom(
                AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams.to_proto(
                    resource.reduce_fraction_less_than_params
                )
            )
        else:
            res.ClearField("reduce_fraction_less_than_params")
        if AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams.to_proto(
            resource.reduce_make_distribution_params
        ):
            res.reduce_make_distribution_params.CopyFrom(
                AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams.to_proto(
                    resource.reduce_make_distribution_params
                )
            )
        else:
            res.ClearField("reduce_make_distribution_params")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionThresholdAggregations(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
            reduce_fraction_less_than_params=AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams.from_proto(
                resource.reduce_fraction_less_than_params
            ),
            reduce_make_distribution_params=AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams.from_proto(
                resource.reduce_make_distribution_params
            ),
        )


class AlertPolicyConditionsConditionThresholdAggregationsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionThresholdAggregations.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionThresholdAggregations.from_proto(i)
            for i in resources
        ]


class AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(
    object
):
    def __init__(self, threshold: float = None):
        self.threshold = threshold

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams()
        )
        if Primitive.to_proto(resource.threshold):
            res.threshold = Primitive.to_proto(resource.threshold)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams(
            threshold=Primitive.from_proto(resource.threshold),
        )


class AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionThresholdAggregationsReduceFractionLessThanParams.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(
    object
):
    def __init__(self, bucket_options: dict = None, exemplar_sampling: dict = None):
        self.bucket_options = bucket_options
        self.exemplar_sampling = exemplar_sampling

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams()
        )
        if AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions.to_proto(
            resource.bucket_options
        ):
            res.bucket_options.CopyFrom(
                AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions.to_proto(
                    resource.bucket_options
                )
            )
        else:
            res.ClearField("bucket_options")
        if AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling.to_proto(
            resource.exemplar_sampling
        ):
            res.exemplar_sampling.CopyFrom(
                AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling.to_proto(
                    resource.exemplar_sampling
                )
            )
        else:
            res.ClearField("exemplar_sampling")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams(
            bucket_options=AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions.from_proto(
                resource.bucket_options
            ),
            exemplar_sampling=AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling.from_proto(
                resource.exemplar_sampling
            ),
        )


class AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParams.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(
    object
):
    def __init__(
        self,
        linear_buckets: dict = None,
        exponential_buckets: dict = None,
        explicit_buckets: dict = None,
    ):
        self.linear_buckets = linear_buckets
        self.exponential_buckets = exponential_buckets
        self.explicit_buckets = explicit_buckets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions()
        )
        if AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
            resource.linear_buckets
        ):
            res.linear_buckets.CopyFrom(
                AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                    resource.linear_buckets
                )
            )
        else:
            res.ClearField("linear_buckets")
        if AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
            resource.exponential_buckets
        ):
            res.exponential_buckets.CopyFrom(
                AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                    resource.exponential_buckets
                )
            )
        else:
            res.ClearField("exponential_buckets")
        if AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
            resource.explicit_buckets
        ):
            res.explicit_buckets.CopyFrom(
                AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                    resource.explicit_buckets
                )
            )
        else:
            res.ClearField("explicit_buckets")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions(
            linear_buckets=AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.from_proto(
                resource.linear_buckets
            ),
            exponential_buckets=AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.from_proto(
                resource.exponential_buckets
            ),
            explicit_buckets=AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.from_proto(
                resource.explicit_buckets
            ),
        )


class AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptions.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(
    object
):
    def __init__(
        self, num_finite_buckets: int = None, width: float = None, offset: float = None
    ):
        self.num_finite_buckets = num_finite_buckets
        self.width = width
        self.offset = offset

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.width):
            res.width = Primitive.to_proto(resource.width)
        if Primitive.to_proto(resource.offset):
            res.offset = Primitive.to_proto(resource.offset)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(
            num_finite_buckets=Primitive.from_proto(resource.num_finite_buckets),
            width=Primitive.from_proto(resource.width),
            offset=Primitive.from_proto(resource.offset),
        )


class AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
    object
):
    def __init__(
        self,
        num_finite_buckets: int = None,
        growth_factor: float = None,
        scale: float = None,
    ):
        self.num_finite_buckets = num_finite_buckets
        self.growth_factor = growth_factor
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.growth_factor):
            res.growth_factor = Primitive.to_proto(resource.growth_factor)
        if Primitive.to_proto(resource.scale):
            res.scale = Primitive.to_proto(resource.scale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
            num_finite_buckets=Primitive.from_proto(resource.num_finite_buckets),
            growth_factor=Primitive.from_proto(resource.growth_factor),
            scale=Primitive.from_proto(resource.scale),
        )


class AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
    object
):
    def __init__(self, bounds: list = None):
        self.bounds = bounds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets()
        )
        if float64Array.to_proto(resource.bounds):
            res.bounds.extend(float64Array.to_proto(resource.bounds))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
            bounds=float64Array.from_proto(resource.bounds),
        )


class AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(
    object
):
    def __init__(self, minimum_value: float = None):
        self.minimum_value = minimum_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling()
        )
        if Primitive.to_proto(resource.minimum_value):
            res.minimum_value = Primitive.to_proto(resource.minimum_value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling(
            minimum_value=Primitive.from_proto(resource.minimum_value),
        )


class AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSamplingArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionThresholdAggregationsReduceMakeDistributionParamsExemplarSampling.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionThresholdDenominatorAggregations(object):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
        reduce_fraction_less_than_params: dict = None,
        reduce_make_distribution_params: dict = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields
        self.reduce_fraction_less_than_params = reduce_fraction_less_than_params
        self.reduce_make_distribution_params = reduce_make_distribution_params

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregations()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        if AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams.to_proto(
            resource.reduce_fraction_less_than_params
        ):
            res.reduce_fraction_less_than_params.CopyFrom(
                AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams.to_proto(
                    resource.reduce_fraction_less_than_params
                )
            )
        else:
            res.ClearField("reduce_fraction_less_than_params")
        if AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams.to_proto(
            resource.reduce_make_distribution_params
        ):
            res.reduce_make_distribution_params.CopyFrom(
                AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams.to_proto(
                    resource.reduce_make_distribution_params
                )
            )
        else:
            res.ClearField("reduce_make_distribution_params")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionThresholdDenominatorAggregations(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
            reduce_fraction_less_than_params=AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams.from_proto(
                resource.reduce_fraction_less_than_params
            ),
            reduce_make_distribution_params=AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams.from_proto(
                resource.reduce_make_distribution_params
            ),
        )


class AlertPolicyConditionsConditionThresholdDenominatorAggregationsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionThresholdDenominatorAggregations.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionThresholdDenominatorAggregations.from_proto(i)
            for i in resources
        ]


class AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(
    object
):
    def __init__(self, threshold: float = None):
        self.threshold = threshold

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams()
        )
        if Primitive.to_proto(resource.threshold):
            res.threshold = Primitive.to_proto(resource.threshold)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams(
            threshold=Primitive.from_proto(resource.threshold),
        )


class AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceFractionLessThanParams.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(
    object
):
    def __init__(self, bucket_options: dict = None, exemplar_sampling: dict = None):
        self.bucket_options = bucket_options
        self.exemplar_sampling = exemplar_sampling

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams()
        )
        if AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions.to_proto(
            resource.bucket_options
        ):
            res.bucket_options.CopyFrom(
                AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions.to_proto(
                    resource.bucket_options
                )
            )
        else:
            res.ClearField("bucket_options")
        if AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling.to_proto(
            resource.exemplar_sampling
        ):
            res.exemplar_sampling.CopyFrom(
                AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling.to_proto(
                    resource.exemplar_sampling
                )
            )
        else:
            res.ClearField("exemplar_sampling")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams(
            bucket_options=AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions.from_proto(
                resource.bucket_options
            ),
            exemplar_sampling=AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling.from_proto(
                resource.exemplar_sampling
            ),
        )


class AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParams.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(
    object
):
    def __init__(
        self,
        linear_buckets: dict = None,
        exponential_buckets: dict = None,
        explicit_buckets: dict = None,
    ):
        self.linear_buckets = linear_buckets
        self.exponential_buckets = exponential_buckets
        self.explicit_buckets = explicit_buckets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions()
        )
        if AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
            resource.linear_buckets
        ):
            res.linear_buckets.CopyFrom(
                AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                    resource.linear_buckets
                )
            )
        else:
            res.ClearField("linear_buckets")
        if AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
            resource.exponential_buckets
        ):
            res.exponential_buckets.CopyFrom(
                AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                    resource.exponential_buckets
                )
            )
        else:
            res.ClearField("exponential_buckets")
        if AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
            resource.explicit_buckets
        ):
            res.explicit_buckets.CopyFrom(
                AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                    resource.explicit_buckets
                )
            )
        else:
            res.ClearField("explicit_buckets")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions(
            linear_buckets=AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.from_proto(
                resource.linear_buckets
            ),
            exponential_buckets=AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.from_proto(
                resource.exponential_buckets
            ),
            explicit_buckets=AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.from_proto(
                resource.explicit_buckets
            ),
        )


class AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptions.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(
    object
):
    def __init__(
        self, num_finite_buckets: int = None, width: float = None, offset: float = None
    ):
        self.num_finite_buckets = num_finite_buckets
        self.width = width
        self.offset = offset

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.width):
            res.width = Primitive.to_proto(resource.width)
        if Primitive.to_proto(resource.offset):
            res.offset = Primitive.to_proto(resource.offset)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(
            num_finite_buckets=Primitive.from_proto(resource.num_finite_buckets),
            width=Primitive.from_proto(resource.width),
            offset=Primitive.from_proto(resource.offset),
        )


class AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
    object
):
    def __init__(
        self,
        num_finite_buckets: int = None,
        growth_factor: float = None,
        scale: float = None,
    ):
        self.num_finite_buckets = num_finite_buckets
        self.growth_factor = growth_factor
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.growth_factor):
            res.growth_factor = Primitive.to_proto(resource.growth_factor)
        if Primitive.to_proto(resource.scale):
            res.scale = Primitive.to_proto(resource.scale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
            num_finite_buckets=Primitive.from_proto(resource.num_finite_buckets),
            growth_factor=Primitive.from_proto(resource.growth_factor),
            scale=Primitive.from_proto(resource.scale),
        )


class AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
    object
):
    def __init__(self, bounds: list = None):
        self.bounds = bounds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets()
        )
        if float64Array.to_proto(resource.bounds):
            res.bounds.extend(float64Array.to_proto(resource.bounds))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
            bounds=float64Array.from_proto(resource.bounds),
        )


class AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(
    object
):
    def __init__(self, minimum_value: float = None):
        self.minimum_value = minimum_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling()
        )
        if Primitive.to_proto(resource.minimum_value):
            res.minimum_value = Primitive.to_proto(resource.minimum_value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling(
            minimum_value=Primitive.from_proto(resource.minimum_value),
        )


class AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSamplingArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionThresholdDenominatorAggregationsReduceMakeDistributionParamsExemplarSampling.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionThresholdTrigger(object):
    def __init__(self, count: int = None, percent: float = None):
        self.count = count
        self.percent = percent

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdTrigger()
        )
        if Primitive.to_proto(resource.count):
            res.count = Primitive.to_proto(resource.count)
        if Primitive.to_proto(resource.percent):
            res.percent = Primitive.to_proto(resource.percent)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionThresholdTrigger(
            count=Primitive.from_proto(resource.count),
            percent=Primitive.from_proto(resource.percent),
        )


class AlertPolicyConditionsConditionThresholdTriggerArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionThresholdTrigger.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionThresholdTrigger.from_proto(i)
            for i in resources
        ]


class AlertPolicyConditionsConditionAbsent(object):
    def __init__(
        self,
        filter: str = None,
        aggregations: list = None,
        duration: dict = None,
        trigger: dict = None,
    ):
        self.filter = filter
        self.aggregations = aggregations
        self.duration = duration
        self.trigger = trigger

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyConditionsConditionAbsent()
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if AlertPolicyConditionsConditionAbsentAggregationsArray.to_proto(
            resource.aggregations
        ):
            res.aggregations.extend(
                AlertPolicyConditionsConditionAbsentAggregationsArray.to_proto(
                    resource.aggregations
                )
            )
        if AlertPolicyConditionsConditionAbsentDuration.to_proto(resource.duration):
            res.duration.CopyFrom(
                AlertPolicyConditionsConditionAbsentDuration.to_proto(resource.duration)
            )
        else:
            res.ClearField("duration")
        if AlertPolicyConditionsConditionAbsentTrigger.to_proto(resource.trigger):
            res.trigger.CopyFrom(
                AlertPolicyConditionsConditionAbsentTrigger.to_proto(resource.trigger)
            )
        else:
            res.ClearField("trigger")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionAbsent(
            filter=Primitive.from_proto(resource.filter),
            aggregations=AlertPolicyConditionsConditionAbsentAggregationsArray.from_proto(
                resource.aggregations
            ),
            duration=AlertPolicyConditionsConditionAbsentDuration.from_proto(
                resource.duration
            ),
            trigger=AlertPolicyConditionsConditionAbsentTrigger.from_proto(
                resource.trigger
            ),
        )


class AlertPolicyConditionsConditionAbsentArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AlertPolicyConditionsConditionAbsent.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AlertPolicyConditionsConditionAbsent.from_proto(i) for i in resources]


class AlertPolicyConditionsConditionAbsentAggregations(object):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
        reduce_fraction_less_than_params: dict = None,
        reduce_make_distribution_params: dict = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields
        self.reduce_fraction_less_than_params = reduce_fraction_less_than_params
        self.reduce_make_distribution_params = reduce_make_distribution_params

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionAbsentAggregations()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        if AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams.to_proto(
            resource.reduce_fraction_less_than_params
        ):
            res.reduce_fraction_less_than_params.CopyFrom(
                AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams.to_proto(
                    resource.reduce_fraction_less_than_params
                )
            )
        else:
            res.ClearField("reduce_fraction_less_than_params")
        if AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams.to_proto(
            resource.reduce_make_distribution_params
        ):
            res.reduce_make_distribution_params.CopyFrom(
                AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams.to_proto(
                    resource.reduce_make_distribution_params
                )
            )
        else:
            res.ClearField("reduce_make_distribution_params")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionAbsentAggregations(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
            reduce_fraction_less_than_params=AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams.from_proto(
                resource.reduce_fraction_less_than_params
            ),
            reduce_make_distribution_params=AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams.from_proto(
                resource.reduce_make_distribution_params
            ),
        )


class AlertPolicyConditionsConditionAbsentAggregationsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionAbsentAggregations.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionAbsentAggregations.from_proto(i)
            for i in resources
        ]


class AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(
    object
):
    def __init__(self, threshold: float = None):
        self.threshold = threshold

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams()
        )
        if Primitive.to_proto(resource.threshold):
            res.threshold = Primitive.to_proto(resource.threshold)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams(
            threshold=Primitive.from_proto(resource.threshold),
        )


class AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionAbsentAggregationsReduceFractionLessThanParams.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(
    object
):
    def __init__(self, bucket_options: dict = None, exemplar_sampling: dict = None):
        self.bucket_options = bucket_options
        self.exemplar_sampling = exemplar_sampling

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams()
        )
        if AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions.to_proto(
            resource.bucket_options
        ):
            res.bucket_options.CopyFrom(
                AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions.to_proto(
                    resource.bucket_options
                )
            )
        else:
            res.ClearField("bucket_options")
        if AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling.to_proto(
            resource.exemplar_sampling
        ):
            res.exemplar_sampling.CopyFrom(
                AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling.to_proto(
                    resource.exemplar_sampling
                )
            )
        else:
            res.ClearField("exemplar_sampling")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams(
            bucket_options=AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions.from_proto(
                resource.bucket_options
            ),
            exemplar_sampling=AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling.from_proto(
                resource.exemplar_sampling
            ),
        )


class AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParams.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(
    object
):
    def __init__(
        self,
        linear_buckets: dict = None,
        exponential_buckets: dict = None,
        explicit_buckets: dict = None,
    ):
        self.linear_buckets = linear_buckets
        self.exponential_buckets = exponential_buckets
        self.explicit_buckets = explicit_buckets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions()
        )
        if AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
            resource.linear_buckets
        ):
            res.linear_buckets.CopyFrom(
                AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                    resource.linear_buckets
                )
            )
        else:
            res.ClearField("linear_buckets")
        if AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
            resource.exponential_buckets
        ):
            res.exponential_buckets.CopyFrom(
                AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                    resource.exponential_buckets
                )
            )
        else:
            res.ClearField("exponential_buckets")
        if AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
            resource.explicit_buckets
        ):
            res.explicit_buckets.CopyFrom(
                AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                    resource.explicit_buckets
                )
            )
        else:
            res.ClearField("explicit_buckets")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions(
            linear_buckets=AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.from_proto(
                resource.linear_buckets
            ),
            exponential_buckets=AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.from_proto(
                resource.exponential_buckets
            ),
            explicit_buckets=AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.from_proto(
                resource.explicit_buckets
            ),
        )


class AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptions.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(
    object
):
    def __init__(
        self, num_finite_buckets: int = None, width: float = None, offset: float = None
    ):
        self.num_finite_buckets = num_finite_buckets
        self.width = width
        self.offset = offset

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.width):
            res.width = Primitive.to_proto(resource.width)
        if Primitive.to_proto(resource.offset):
            res.offset = Primitive.to_proto(resource.offset)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(
            num_finite_buckets=Primitive.from_proto(resource.num_finite_buckets),
            width=Primitive.from_proto(resource.width),
            offset=Primitive.from_proto(resource.offset),
        )


class AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
    object
):
    def __init__(
        self,
        num_finite_buckets: int = None,
        growth_factor: float = None,
        scale: float = None,
    ):
        self.num_finite_buckets = num_finite_buckets
        self.growth_factor = growth_factor
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.growth_factor):
            res.growth_factor = Primitive.to_proto(resource.growth_factor)
        if Primitive.to_proto(resource.scale):
            res.scale = Primitive.to_proto(resource.scale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
            num_finite_buckets=Primitive.from_proto(resource.num_finite_buckets),
            growth_factor=Primitive.from_proto(resource.growth_factor),
            scale=Primitive.from_proto(resource.scale),
        )


class AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
    object
):
    def __init__(self, bounds: list = None):
        self.bounds = bounds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets()
        )
        if float64Array.to_proto(resource.bounds):
            res.bounds.extend(float64Array.to_proto(resource.bounds))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
            bounds=float64Array.from_proto(resource.bounds),
        )


class AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(
    object
):
    def __init__(self, minimum_value: float = None):
        self.minimum_value = minimum_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling()
        )
        if Primitive.to_proto(resource.minimum_value):
            res.minimum_value = Primitive.to_proto(resource.minimum_value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling(
            minimum_value=Primitive.from_proto(resource.minimum_value),
        )


class AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSamplingArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionAbsentAggregationsReduceMakeDistributionParamsExemplarSampling.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionAbsentDuration(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyConditionsConditionAbsentDuration()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionAbsentDuration(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class AlertPolicyConditionsConditionAbsentDurationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionAbsentDuration.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionAbsentDuration.from_proto(i)
            for i in resources
        ]


class AlertPolicyConditionsConditionAbsentTrigger(object):
    def __init__(self, count: int = None, percent: float = None):
        self.count = count
        self.percent = percent

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyConditionsConditionAbsentTrigger()
        if Primitive.to_proto(resource.count):
            res.count = Primitive.to_proto(resource.count)
        if Primitive.to_proto(resource.percent):
            res.percent = Primitive.to_proto(resource.percent)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionAbsentTrigger(
            count=Primitive.from_proto(resource.count),
            percent=Primitive.from_proto(resource.percent),
        )


class AlertPolicyConditionsConditionAbsentTriggerArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionAbsentTrigger.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionAbsentTrigger.from_proto(i) for i in resources
        ]


class AlertPolicyConditionsConditionMatchedLog(object):
    def __init__(self, filter: str = None, label_extractors: dict = None):
        self.filter = filter
        self.label_extractors = label_extractors

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyConditionsConditionMatchedLog()
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if Primitive.to_proto(resource.label_extractors):
            res.label_extractors = Primitive.to_proto(resource.label_extractors)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionMatchedLog(
            filter=Primitive.from_proto(resource.filter),
            label_extractors=Primitive.from_proto(resource.label_extractors),
        )


class AlertPolicyConditionsConditionMatchedLogArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AlertPolicyConditionsConditionMatchedLog.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionMatchedLog.from_proto(i) for i in resources
        ]


class AlertPolicyConditionsConditionClusterOutlier(object):
    def __init__(self, filter: str = None):
        self.filter = filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyConditionsConditionClusterOutlier()
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionClusterOutlier(
            filter=Primitive.from_proto(resource.filter),
        )


class AlertPolicyConditionsConditionClusterOutlierArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionClusterOutlier.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionClusterOutlier.from_proto(i)
            for i in resources
        ]


class AlertPolicyConditionsConditionRate(object):
    def __init__(
        self,
        filter: str = None,
        aggregations: list = None,
        comparison: str = None,
        threshold_value: float = None,
        time_window: dict = None,
        trigger: dict = None,
    ):
        self.filter = filter
        self.aggregations = aggregations
        self.comparison = comparison
        self.threshold_value = threshold_value
        self.time_window = time_window
        self.trigger = trigger

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyConditionsConditionRate()
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if AlertPolicyConditionsConditionRateAggregationsArray.to_proto(
            resource.aggregations
        ):
            res.aggregations.extend(
                AlertPolicyConditionsConditionRateAggregationsArray.to_proto(
                    resource.aggregations
                )
            )
        if AlertPolicyConditionsConditionRateComparisonEnum.to_proto(
            resource.comparison
        ):
            res.comparison = AlertPolicyConditionsConditionRateComparisonEnum.to_proto(
                resource.comparison
            )
        if Primitive.to_proto(resource.threshold_value):
            res.threshold_value = Primitive.to_proto(resource.threshold_value)
        if AlertPolicyConditionsConditionRateTimeWindow.to_proto(resource.time_window):
            res.time_window.CopyFrom(
                AlertPolicyConditionsConditionRateTimeWindow.to_proto(
                    resource.time_window
                )
            )
        else:
            res.ClearField("time_window")
        if AlertPolicyConditionsConditionRateTrigger.to_proto(resource.trigger):
            res.trigger.CopyFrom(
                AlertPolicyConditionsConditionRateTrigger.to_proto(resource.trigger)
            )
        else:
            res.ClearField("trigger")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionRate(
            filter=Primitive.from_proto(resource.filter),
            aggregations=AlertPolicyConditionsConditionRateAggregationsArray.from_proto(
                resource.aggregations
            ),
            comparison=AlertPolicyConditionsConditionRateComparisonEnum.from_proto(
                resource.comparison
            ),
            threshold_value=Primitive.from_proto(resource.threshold_value),
            time_window=AlertPolicyConditionsConditionRateTimeWindow.from_proto(
                resource.time_window
            ),
            trigger=AlertPolicyConditionsConditionRateTrigger.from_proto(
                resource.trigger
            ),
        )


class AlertPolicyConditionsConditionRateArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AlertPolicyConditionsConditionRate.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AlertPolicyConditionsConditionRate.from_proto(i) for i in resources]


class AlertPolicyConditionsConditionRateAggregations(object):
    def __init__(
        self,
        alignment_period: str = None,
        per_series_aligner: str = None,
        cross_series_reducer: str = None,
        group_by_fields: list = None,
        reduce_fraction_less_than_params: dict = None,
        reduce_make_distribution_params: dict = None,
    ):
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.cross_series_reducer = cross_series_reducer
        self.group_by_fields = group_by_fields
        self.reduce_fraction_less_than_params = reduce_fraction_less_than_params
        self.reduce_make_distribution_params = reduce_make_distribution_params

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionRateAggregations()
        )
        if Primitive.to_proto(resource.alignment_period):
            res.alignment_period = Primitive.to_proto(resource.alignment_period)
        if AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum.to_proto(
            resource.per_series_aligner
        ):
            res.per_series_aligner = AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum.to_proto(
                resource.per_series_aligner
            )
        if AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum.to_proto(
            resource.cross_series_reducer
        ):
            res.cross_series_reducer = AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum.to_proto(
                resource.cross_series_reducer
            )
        if Primitive.to_proto(resource.group_by_fields):
            res.group_by_fields.extend(Primitive.to_proto(resource.group_by_fields))
        if AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams.to_proto(
            resource.reduce_fraction_less_than_params
        ):
            res.reduce_fraction_less_than_params.CopyFrom(
                AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams.to_proto(
                    resource.reduce_fraction_less_than_params
                )
            )
        else:
            res.ClearField("reduce_fraction_less_than_params")
        if AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams.to_proto(
            resource.reduce_make_distribution_params
        ):
            res.reduce_make_distribution_params.CopyFrom(
                AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams.to_proto(
                    resource.reduce_make_distribution_params
                )
            )
        else:
            res.ClearField("reduce_make_distribution_params")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionRateAggregations(
            alignment_period=Primitive.from_proto(resource.alignment_period),
            per_series_aligner=AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum.from_proto(
                resource.per_series_aligner
            ),
            cross_series_reducer=AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum.from_proto(
                resource.cross_series_reducer
            ),
            group_by_fields=Primitive.from_proto(resource.group_by_fields),
            reduce_fraction_less_than_params=AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams.from_proto(
                resource.reduce_fraction_less_than_params
            ),
            reduce_make_distribution_params=AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams.from_proto(
                resource.reduce_make_distribution_params
            ),
        )


class AlertPolicyConditionsConditionRateAggregationsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionRateAggregations.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionRateAggregations.from_proto(i)
            for i in resources
        ]


class AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(
    object
):
    def __init__(self, threshold: float = None):
        self.threshold = threshold

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams()
        )
        if Primitive.to_proto(resource.threshold):
            res.threshold = Primitive.to_proto(resource.threshold)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams(
            threshold=Primitive.from_proto(resource.threshold),
        )


class AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionRateAggregationsReduceFractionLessThanParams.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(
    object
):
    def __init__(self, bucket_options: dict = None, exemplar_sampling: dict = None):
        self.bucket_options = bucket_options
        self.exemplar_sampling = exemplar_sampling

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams()
        )
        if AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions.to_proto(
            resource.bucket_options
        ):
            res.bucket_options.CopyFrom(
                AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions.to_proto(
                    resource.bucket_options
                )
            )
        else:
            res.ClearField("bucket_options")
        if AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling.to_proto(
            resource.exemplar_sampling
        ):
            res.exemplar_sampling.CopyFrom(
                AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling.to_proto(
                    resource.exemplar_sampling
                )
            )
        else:
            res.ClearField("exemplar_sampling")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams(
            bucket_options=AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions.from_proto(
                resource.bucket_options
            ),
            exemplar_sampling=AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling.from_proto(
                resource.exemplar_sampling
            ),
        )


class AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParams.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(
    object
):
    def __init__(
        self,
        linear_buckets: dict = None,
        exponential_buckets: dict = None,
        explicit_buckets: dict = None,
    ):
        self.linear_buckets = linear_buckets
        self.exponential_buckets = exponential_buckets
        self.explicit_buckets = explicit_buckets

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions()
        )
        if AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
            resource.linear_buckets
        ):
            res.linear_buckets.CopyFrom(
                AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                    resource.linear_buckets
                )
            )
        else:
            res.ClearField("linear_buckets")
        if AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
            resource.exponential_buckets
        ):
            res.exponential_buckets.CopyFrom(
                AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                    resource.exponential_buckets
                )
            )
        else:
            res.ClearField("exponential_buckets")
        if AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
            resource.explicit_buckets
        ):
            res.explicit_buckets.CopyFrom(
                AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                    resource.explicit_buckets
                )
            )
        else:
            res.ClearField("explicit_buckets")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions(
            linear_buckets=AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.from_proto(
                resource.linear_buckets
            ),
            exponential_buckets=AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.from_proto(
                resource.exponential_buckets
            ),
            explicit_buckets=AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.from_proto(
                resource.explicit_buckets
            ),
        )


class AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptions.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(
    object
):
    def __init__(
        self, num_finite_buckets: int = None, width: float = None, offset: float = None
    ):
        self.num_finite_buckets = num_finite_buckets
        self.width = width
        self.offset = offset

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.width):
            res.width = Primitive.to_proto(resource.width)
        if Primitive.to_proto(resource.offset):
            res.offset = Primitive.to_proto(resource.offset)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets(
            num_finite_buckets=Primitive.from_proto(resource.num_finite_buckets),
            width=Primitive.from_proto(resource.width),
            offset=Primitive.from_proto(resource.offset),
        )


class AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsLinearBuckets.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
    object
):
    def __init__(
        self,
        num_finite_buckets: int = None,
        growth_factor: float = None,
        scale: float = None,
    ):
        self.num_finite_buckets = num_finite_buckets
        self.growth_factor = growth_factor
        self.scale = scale

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets()
        )
        if Primitive.to_proto(resource.num_finite_buckets):
            res.num_finite_buckets = Primitive.to_proto(resource.num_finite_buckets)
        if Primitive.to_proto(resource.growth_factor):
            res.growth_factor = Primitive.to_proto(resource.growth_factor)
        if Primitive.to_proto(resource.scale):
            res.scale = Primitive.to_proto(resource.scale)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets(
            num_finite_buckets=Primitive.from_proto(resource.num_finite_buckets),
            growth_factor=Primitive.from_proto(resource.growth_factor),
            scale=Primitive.from_proto(resource.scale),
        )


class AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExponentialBuckets.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
    object
):
    def __init__(self, bounds: list = None):
        self.bounds = bounds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets()
        )
        if float64Array.to_proto(resource.bounds):
            res.bounds.extend(float64Array.to_proto(resource.bounds))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets(
            bounds=float64Array.from_proto(resource.bounds),
        )


class AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBucketsArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsBucketOptionsExplicitBuckets.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(
    object
):
    def __init__(self, minimum_value: float = None):
        self.minimum_value = minimum_value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling()
        )
        if Primitive.to_proto(resource.minimum_value):
            res.minimum_value = Primitive.to_proto(resource.minimum_value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling(
            minimum_value=Primitive.from_proto(resource.minimum_value),
        )


class AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSamplingArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionRateAggregationsReduceMakeDistributionParamsExemplarSampling.from_proto(
                i
            )
            for i in resources
        ]


class AlertPolicyConditionsConditionRateTimeWindow(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyConditionsConditionRateTimeWindow()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionRateTimeWindow(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class AlertPolicyConditionsConditionRateTimeWindowArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionRateTimeWindow.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionRateTimeWindow.from_proto(i)
            for i in resources
        ]


class AlertPolicyConditionsConditionRateTrigger(object):
    def __init__(self, count: int = None, percent: float = None):
        self.count = count
        self.percent = percent

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyConditionsConditionRateTrigger()
        if Primitive.to_proto(resource.count):
            res.count = Primitive.to_proto(resource.count)
        if Primitive.to_proto(resource.percent):
            res.percent = Primitive.to_proto(resource.percent)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionRateTrigger(
            count=Primitive.from_proto(resource.count),
            percent=Primitive.from_proto(resource.percent),
        )


class AlertPolicyConditionsConditionRateTriggerArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionRateTrigger.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionRateTrigger.from_proto(i) for i in resources
        ]


class AlertPolicyConditionsConditionUpMon(object):
    def __init__(
        self,
        filter: str = None,
        endpoint_id: str = None,
        check_id: str = None,
        duration: dict = None,
        trigger: dict = None,
    ):
        self.filter = filter
        self.endpoint_id = endpoint_id
        self.check_id = check_id
        self.duration = duration
        self.trigger = trigger

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyConditionsConditionUpMon()
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if Primitive.to_proto(resource.endpoint_id):
            res.endpoint_id = Primitive.to_proto(resource.endpoint_id)
        if Primitive.to_proto(resource.check_id):
            res.check_id = Primitive.to_proto(resource.check_id)
        if AlertPolicyConditionsConditionUpMonDuration.to_proto(resource.duration):
            res.duration.CopyFrom(
                AlertPolicyConditionsConditionUpMonDuration.to_proto(resource.duration)
            )
        else:
            res.ClearField("duration")
        if AlertPolicyConditionsConditionUpMonTrigger.to_proto(resource.trigger):
            res.trigger.CopyFrom(
                AlertPolicyConditionsConditionUpMonTrigger.to_proto(resource.trigger)
            )
        else:
            res.ClearField("trigger")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionUpMon(
            filter=Primitive.from_proto(resource.filter),
            endpoint_id=Primitive.from_proto(resource.endpoint_id),
            check_id=Primitive.from_proto(resource.check_id),
            duration=AlertPolicyConditionsConditionUpMonDuration.from_proto(
                resource.duration
            ),
            trigger=AlertPolicyConditionsConditionUpMonTrigger.from_proto(
                resource.trigger
            ),
        )


class AlertPolicyConditionsConditionUpMonArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AlertPolicyConditionsConditionUpMon.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AlertPolicyConditionsConditionUpMon.from_proto(i) for i in resources]


class AlertPolicyConditionsConditionUpMonDuration(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyConditionsConditionUpMonDuration()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionUpMonDuration(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class AlertPolicyConditionsConditionUpMonDurationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionUpMonDuration.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionUpMonDuration.from_proto(i) for i in resources
        ]


class AlertPolicyConditionsConditionUpMonTrigger(object):
    def __init__(self, count: int = None, percent: float = None):
        self.count = count
        self.percent = percent

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyConditionsConditionUpMonTrigger()
        if Primitive.to_proto(resource.count):
            res.count = Primitive.to_proto(resource.count)
        if Primitive.to_proto(resource.percent):
            res.percent = Primitive.to_proto(resource.percent)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionUpMonTrigger(
            count=Primitive.from_proto(resource.count),
            percent=Primitive.from_proto(resource.percent),
        )


class AlertPolicyConditionsConditionUpMonTriggerArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionUpMonTrigger.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionUpMonTrigger.from_proto(i) for i in resources
        ]


class AlertPolicyConditionsConditionProcessCount(object):
    def __init__(
        self,
        process: str = None,
        user: str = None,
        filter: str = None,
        comparison: str = None,
        process_count_threshold: int = None,
        trigger: dict = None,
        duration: dict = None,
    ):
        self.process = process
        self.user = user
        self.filter = filter
        self.comparison = comparison
        self.process_count_threshold = process_count_threshold
        self.trigger = trigger
        self.duration = duration

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyConditionsConditionProcessCount()
        if Primitive.to_proto(resource.process):
            res.process = Primitive.to_proto(resource.process)
        if Primitive.to_proto(resource.user):
            res.user = Primitive.to_proto(resource.user)
        if Primitive.to_proto(resource.filter):
            res.filter = Primitive.to_proto(resource.filter)
        if AlertPolicyConditionsConditionProcessCountComparisonEnum.to_proto(
            resource.comparison
        ):
            res.comparison = AlertPolicyConditionsConditionProcessCountComparisonEnum.to_proto(
                resource.comparison
            )
        if Primitive.to_proto(resource.process_count_threshold):
            res.process_count_threshold = Primitive.to_proto(
                resource.process_count_threshold
            )
        if AlertPolicyConditionsConditionProcessCountTrigger.to_proto(resource.trigger):
            res.trigger.CopyFrom(
                AlertPolicyConditionsConditionProcessCountTrigger.to_proto(
                    resource.trigger
                )
            )
        else:
            res.ClearField("trigger")
        if AlertPolicyConditionsConditionProcessCountDuration.to_proto(
            resource.duration
        ):
            res.duration.CopyFrom(
                AlertPolicyConditionsConditionProcessCountDuration.to_proto(
                    resource.duration
                )
            )
        else:
            res.ClearField("duration")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionProcessCount(
            process=Primitive.from_proto(resource.process),
            user=Primitive.from_proto(resource.user),
            filter=Primitive.from_proto(resource.filter),
            comparison=AlertPolicyConditionsConditionProcessCountComparisonEnum.from_proto(
                resource.comparison
            ),
            process_count_threshold=Primitive.from_proto(
                resource.process_count_threshold
            ),
            trigger=AlertPolicyConditionsConditionProcessCountTrigger.from_proto(
                resource.trigger
            ),
            duration=AlertPolicyConditionsConditionProcessCountDuration.from_proto(
                resource.duration
            ),
        )


class AlertPolicyConditionsConditionProcessCountArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionProcessCount.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionProcessCount.from_proto(i) for i in resources
        ]


class AlertPolicyConditionsConditionProcessCountTrigger(object):
    def __init__(self, count: int = None, percent: float = None):
        self.count = count
        self.percent = percent

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionProcessCountTrigger()
        )
        if Primitive.to_proto(resource.count):
            res.count = Primitive.to_proto(resource.count)
        if Primitive.to_proto(resource.percent):
            res.percent = Primitive.to_proto(resource.percent)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionProcessCountTrigger(
            count=Primitive.from_proto(resource.count),
            percent=Primitive.from_proto(resource.percent),
        )


class AlertPolicyConditionsConditionProcessCountTriggerArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionProcessCountTrigger.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionProcessCountTrigger.from_proto(i)
            for i in resources
        ]


class AlertPolicyConditionsConditionProcessCountDuration(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionProcessCountDuration()
        )
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionProcessCountDuration(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class AlertPolicyConditionsConditionProcessCountDurationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionProcessCountDuration.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionProcessCountDuration.from_proto(i)
            for i in resources
        ]


class AlertPolicyConditionsConditionTimeSeriesQueryLanguage(object):
    def __init__(self, query: str = None, summary: str = None):
        self.query = query
        self.summary = summary

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionTimeSeriesQueryLanguage()
        )
        if Primitive.to_proto(resource.query):
            res.query = Primitive.to_proto(resource.query)
        if Primitive.to_proto(resource.summary):
            res.summary = Primitive.to_proto(resource.summary)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionTimeSeriesQueryLanguage(
            query=Primitive.from_proto(resource.query),
            summary=Primitive.from_proto(resource.summary),
        )


class AlertPolicyConditionsConditionTimeSeriesQueryLanguageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionTimeSeriesQueryLanguage.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionTimeSeriesQueryLanguage.from_proto(i)
            for i in resources
        ]


class AlertPolicyConditionsConditionMonitoringQueryLanguage(object):
    def __init__(self, query: str = None, duration: dict = None, trigger: dict = None):
        self.query = query
        self.duration = duration
        self.trigger = trigger

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionMonitoringQueryLanguage()
        )
        if Primitive.to_proto(resource.query):
            res.query = Primitive.to_proto(resource.query)
        if AlertPolicyConditionsConditionMonitoringQueryLanguageDuration.to_proto(
            resource.duration
        ):
            res.duration.CopyFrom(
                AlertPolicyConditionsConditionMonitoringQueryLanguageDuration.to_proto(
                    resource.duration
                )
            )
        else:
            res.ClearField("duration")
        if AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger.to_proto(
            resource.trigger
        ):
            res.trigger.CopyFrom(
                AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger.to_proto(
                    resource.trigger
                )
            )
        else:
            res.ClearField("trigger")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionMonitoringQueryLanguage(
            query=Primitive.from_proto(resource.query),
            duration=AlertPolicyConditionsConditionMonitoringQueryLanguageDuration.from_proto(
                resource.duration
            ),
            trigger=AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger.from_proto(
                resource.trigger
            ),
        )


class AlertPolicyConditionsConditionMonitoringQueryLanguageArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionMonitoringQueryLanguage.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionMonitoringQueryLanguage.from_proto(i)
            for i in resources
        ]


class AlertPolicyConditionsConditionMonitoringQueryLanguageDuration(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageDuration()
        )
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionMonitoringQueryLanguageDuration(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class AlertPolicyConditionsConditionMonitoringQueryLanguageDurationArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionMonitoringQueryLanguageDuration.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionMonitoringQueryLanguageDuration.from_proto(i)
            for i in resources
        ]


class AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(object):
    def __init__(self, count: int = None, percent: float = None):
        self.count = count
        self.percent = percent

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            alert_policy_pb2.MonitoringAlertPolicyConditionsConditionMonitoringQueryLanguageTrigger()
        )
        if Primitive.to_proto(resource.count):
            res.count = Primitive.to_proto(resource.count)
        if Primitive.to_proto(resource.percent):
            res.percent = Primitive.to_proto(resource.percent)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger(
            count=Primitive.from_proto(resource.count),
            percent=Primitive.from_proto(resource.percent),
        )


class AlertPolicyConditionsConditionMonitoringQueryLanguageTriggerArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            AlertPolicyConditionsConditionMonitoringQueryLanguageTrigger.from_proto(i)
            for i in resources
        ]


class AlertPolicyEnabled(object):
    def __init__(self, value: bool = None):
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyEnabled()
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyEnabled(value=Primitive.from_proto(resource.value),)


class AlertPolicyEnabledArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AlertPolicyEnabled.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AlertPolicyEnabled.from_proto(i) for i in resources]


class AlertPolicyValidity(object):
    def __init__(self, code: int = None, message: str = None, details: list = None):
        self.code = code
        self.message = message
        self.details = details

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyValidity()
        if Primitive.to_proto(resource.code):
            res.code = Primitive.to_proto(resource.code)
        if Primitive.to_proto(resource.message):
            res.message = Primitive.to_proto(resource.message)
        if AlertPolicyValidityDetailsArray.to_proto(resource.details):
            res.details.extend(
                AlertPolicyValidityDetailsArray.to_proto(resource.details)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyValidity(
            code=Primitive.from_proto(resource.code),
            message=Primitive.from_proto(resource.message),
            details=AlertPolicyValidityDetailsArray.from_proto(resource.details),
        )


class AlertPolicyValidityArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AlertPolicyValidity.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AlertPolicyValidity.from_proto(i) for i in resources]


class AlertPolicyValidityDetails(object):
    def __init__(self, type_url: str = None, value: str = None):
        self.type_url = type_url
        self.value = value

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyValidityDetails()
        if Primitive.to_proto(resource.type_url):
            res.type_url = Primitive.to_proto(resource.type_url)
        if Primitive.to_proto(resource.value):
            res.value = Primitive.to_proto(resource.value)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyValidityDetails(
            type_url=Primitive.from_proto(resource.type_url),
            value=Primitive.from_proto(resource.value),
        )


class AlertPolicyValidityDetailsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AlertPolicyValidityDetails.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AlertPolicyValidityDetails.from_proto(i) for i in resources]


class AlertPolicyCreationRecord(object):
    def __init__(self, mutate_time: dict = None, mutated_by: str = None):
        self.mutate_time = mutate_time
        self.mutated_by = mutated_by

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyCreationRecord()
        if AlertPolicyCreationRecordMutateTime.to_proto(resource.mutate_time):
            res.mutate_time.CopyFrom(
                AlertPolicyCreationRecordMutateTime.to_proto(resource.mutate_time)
            )
        else:
            res.ClearField("mutate_time")
        if Primitive.to_proto(resource.mutated_by):
            res.mutated_by = Primitive.to_proto(resource.mutated_by)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyCreationRecord(
            mutate_time=AlertPolicyCreationRecordMutateTime.from_proto(
                resource.mutate_time
            ),
            mutated_by=Primitive.from_proto(resource.mutated_by),
        )


class AlertPolicyCreationRecordArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AlertPolicyCreationRecord.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AlertPolicyCreationRecord.from_proto(i) for i in resources]


class AlertPolicyCreationRecordMutateTime(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyCreationRecordMutateTime()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyCreationRecordMutateTime(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class AlertPolicyCreationRecordMutateTimeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AlertPolicyCreationRecordMutateTime.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AlertPolicyCreationRecordMutateTime.from_proto(i) for i in resources]


class AlertPolicyMutationRecord(object):
    def __init__(self, mutate_time: dict = None, mutated_by: str = None):
        self.mutate_time = mutate_time
        self.mutated_by = mutated_by

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyMutationRecord()
        if AlertPolicyMutationRecordMutateTime.to_proto(resource.mutate_time):
            res.mutate_time.CopyFrom(
                AlertPolicyMutationRecordMutateTime.to_proto(resource.mutate_time)
            )
        else:
            res.ClearField("mutate_time")
        if Primitive.to_proto(resource.mutated_by):
            res.mutated_by = Primitive.to_proto(resource.mutated_by)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyMutationRecord(
            mutate_time=AlertPolicyMutationRecordMutateTime.from_proto(
                resource.mutate_time
            ),
            mutated_by=Primitive.from_proto(resource.mutated_by),
        )


class AlertPolicyMutationRecordArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AlertPolicyMutationRecord.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AlertPolicyMutationRecord.from_proto(i) for i in resources]


class AlertPolicyMutationRecordMutateTime(object):
    def __init__(self, seconds: int = None, nanos: int = None):
        self.seconds = seconds
        self.nanos = nanos

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyMutationRecordMutateTime()
        if Primitive.to_proto(resource.seconds):
            res.seconds = Primitive.to_proto(resource.seconds)
        if Primitive.to_proto(resource.nanos):
            res.nanos = Primitive.to_proto(resource.nanos)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyMutationRecordMutateTime(
            seconds=Primitive.from_proto(resource.seconds),
            nanos=Primitive.from_proto(resource.nanos),
        )


class AlertPolicyMutationRecordMutateTimeArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AlertPolicyMutationRecordMutateTime.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AlertPolicyMutationRecordMutateTime.from_proto(i) for i in resources]


class AlertPolicyIncidentStrategy(object):
    def __init__(self, type: str = None):
        self.type = type

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyIncidentStrategy()
        if AlertPolicyIncidentStrategyTypeEnum.to_proto(resource.type):
            res.type = AlertPolicyIncidentStrategyTypeEnum.to_proto(resource.type)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyIncidentStrategy(
            type=AlertPolicyIncidentStrategyTypeEnum.from_proto(resource.type),
        )


class AlertPolicyIncidentStrategyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AlertPolicyIncidentStrategy.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AlertPolicyIncidentStrategy.from_proto(i) for i in resources]


class AlertPolicyMetadata(object):
    def __init__(self, slo_names: list = None):
        self.slo_names = slo_names

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = alert_policy_pb2.MonitoringAlertPolicyMetadata()
        if Primitive.to_proto(resource.slo_names):
            res.slo_names.extend(Primitive.to_proto(resource.slo_names))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return AlertPolicyMetadata(slo_names=Primitive.from_proto(resource.slo_names),)


class AlertPolicyMetadataArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [AlertPolicyMetadata.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [AlertPolicyMetadata.from_proto(i) for i in resources]


class AlertPolicyConditionsResourceStateFilterEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsResourceStateFilterEnum.Value(
            "MonitoringAlertPolicyConditionsResourceStateFilterEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsResourceStateFilterEnum.Name(
            resource
        )[
            len("MonitoringAlertPolicyConditionsResourceStateFilterEnum") :
        ]


class AlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum.Value(
            "MonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringAlertPolicyConditionsConditionThresholdAggregationsPerSeriesAlignerEnum"
            ) :
        ]


class AlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum.Value(
            "MonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringAlertPolicyConditionsConditionThresholdAggregationsCrossSeriesReducerEnum"
            ) :
        ]


class AlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum.Value(
            "MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsPerSeriesAlignerEnum"
            ) :
        ]


class AlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum.Value(
            "MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringAlertPolicyConditionsConditionThresholdDenominatorAggregationsCrossSeriesReducerEnum"
            ) :
        ]


class AlertPolicyConditionsConditionThresholdComparisonEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdComparisonEnum.Value(
            "MonitoringAlertPolicyConditionsConditionThresholdComparisonEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionThresholdComparisonEnum.Name(
            resource
        )[
            len("MonitoringAlertPolicyConditionsConditionThresholdComparisonEnum") :
        ]


class AlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum.Value(
            "MonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringAlertPolicyConditionsConditionAbsentAggregationsPerSeriesAlignerEnum"
            ) :
        ]


class AlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum.Value(
            "MonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringAlertPolicyConditionsConditionAbsentAggregationsCrossSeriesReducerEnum"
            ) :
        ]


class AlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum.Value(
            "MonitoringAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum.Name(
            resource
        )[
            len(
                "MonitoringAlertPolicyConditionsConditionRateAggregationsPerSeriesAlignerEnum"
            ) :
        ]


class AlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum.Value(
            "MonitoringAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum.Name(
            resource
        )[
            len(
                "MonitoringAlertPolicyConditionsConditionRateAggregationsCrossSeriesReducerEnum"
            ) :
        ]


class AlertPolicyConditionsConditionRateComparisonEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionRateComparisonEnum.Value(
            "MonitoringAlertPolicyConditionsConditionRateComparisonEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionRateComparisonEnum.Name(
            resource
        )[
            len("MonitoringAlertPolicyConditionsConditionRateComparisonEnum") :
        ]


class AlertPolicyConditionsConditionProcessCountComparisonEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionProcessCountComparisonEnum.Value(
            "MonitoringAlertPolicyConditionsConditionProcessCountComparisonEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyConditionsConditionProcessCountComparisonEnum.Name(
            resource
        )[
            len("MonitoringAlertPolicyConditionsConditionProcessCountComparisonEnum") :
        ]


class AlertPolicyCombinerEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyCombinerEnum.Value(
            "MonitoringAlertPolicyCombinerEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyCombinerEnum.Name(resource)[
            len("MonitoringAlertPolicyCombinerEnum") :
        ]


class AlertPolicyIncidentStrategyTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyIncidentStrategyTypeEnum.Value(
            "MonitoringAlertPolicyIncidentStrategyTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return alert_policy_pb2.MonitoringAlertPolicyIncidentStrategyTypeEnum.Name(
            resource
        )[len("MonitoringAlertPolicyIncidentStrategyTypeEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
