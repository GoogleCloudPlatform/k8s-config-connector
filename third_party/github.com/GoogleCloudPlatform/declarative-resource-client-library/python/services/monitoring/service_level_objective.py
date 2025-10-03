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
from google3.cloud.graphite.mmv2.services.google.monitoring import (
    service_level_objective_pb2,
)
from google3.cloud.graphite.mmv2.services.google.monitoring import (
    service_level_objective_pb2_grpc,
)

from typing import List


class ServiceLevelObjective(object):
    def __init__(
        self,
        name: str = None,
        display_name: str = None,
        service_level_indicator: dict = None,
        goal: float = None,
        rolling_period: str = None,
        calendar_period: str = None,
        create_time: str = None,
        delete_time: str = None,
        service_management_owned: bool = None,
        user_labels: dict = None,
        project: str = None,
        service: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.display_name = display_name
        self.service_level_indicator = service_level_indicator
        self.goal = goal
        self.rolling_period = rolling_period
        self.calendar_period = calendar_period
        self.user_labels = user_labels
        self.project = project
        self.service = service
        self.service_account_file = service_account_file

    def apply(self):
        stub = (
            service_level_objective_pb2_grpc.MonitoringServiceLevelObjectiveServiceStub(
                channel.Channel()
            )
        )
        request = (
            service_level_objective_pb2.ApplyMonitoringServiceLevelObjectiveRequest()
        )
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if ServiceLevelObjectiveServiceLevelIndicator.to_proto(
            self.service_level_indicator
        ):
            request.resource.service_level_indicator.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicator.to_proto(
                    self.service_level_indicator
                )
            )
        else:
            request.resource.ClearField("service_level_indicator")
        if Primitive.to_proto(self.goal):
            request.resource.goal = Primitive.to_proto(self.goal)

        if Primitive.to_proto(self.rolling_period):
            request.resource.rolling_period = Primitive.to_proto(self.rolling_period)

        if ServiceLevelObjectiveCalendarPeriodEnum.to_proto(self.calendar_period):
            request.resource.calendar_period = (
                ServiceLevelObjectiveCalendarPeriodEnum.to_proto(self.calendar_period)
            )

        if Primitive.to_proto(self.user_labels):
            request.resource.user_labels = Primitive.to_proto(self.user_labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.service):
            request.resource.service = Primitive.to_proto(self.service)

        request.service_account_file = self.service_account_file

        response = stub.ApplyMonitoringServiceLevelObjective(request)
        self.name = Primitive.from_proto(response.name)
        self.display_name = Primitive.from_proto(response.display_name)
        self.service_level_indicator = (
            ServiceLevelObjectiveServiceLevelIndicator.from_proto(
                response.service_level_indicator
            )
        )
        self.goal = Primitive.from_proto(response.goal)
        self.rolling_period = Primitive.from_proto(response.rolling_period)
        self.calendar_period = ServiceLevelObjectiveCalendarPeriodEnum.from_proto(
            response.calendar_period
        )
        self.create_time = Primitive.from_proto(response.create_time)
        self.delete_time = Primitive.from_proto(response.delete_time)
        self.service_management_owned = Primitive.from_proto(
            response.service_management_owned
        )
        self.user_labels = Primitive.from_proto(response.user_labels)
        self.project = Primitive.from_proto(response.project)
        self.service = Primitive.from_proto(response.service)

    def delete(self):
        stub = (
            service_level_objective_pb2_grpc.MonitoringServiceLevelObjectiveServiceStub(
                channel.Channel()
            )
        )
        request = (
            service_level_objective_pb2.DeleteMonitoringServiceLevelObjectiveRequest()
        )
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.display_name):
            request.resource.display_name = Primitive.to_proto(self.display_name)

        if ServiceLevelObjectiveServiceLevelIndicator.to_proto(
            self.service_level_indicator
        ):
            request.resource.service_level_indicator.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicator.to_proto(
                    self.service_level_indicator
                )
            )
        else:
            request.resource.ClearField("service_level_indicator")
        if Primitive.to_proto(self.goal):
            request.resource.goal = Primitive.to_proto(self.goal)

        if Primitive.to_proto(self.rolling_period):
            request.resource.rolling_period = Primitive.to_proto(self.rolling_period)

        if ServiceLevelObjectiveCalendarPeriodEnum.to_proto(self.calendar_period):
            request.resource.calendar_period = (
                ServiceLevelObjectiveCalendarPeriodEnum.to_proto(self.calendar_period)
            )

        if Primitive.to_proto(self.user_labels):
            request.resource.user_labels = Primitive.to_proto(self.user_labels)

        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        if Primitive.to_proto(self.service):
            request.resource.service = Primitive.to_proto(self.service)

        response = stub.DeleteMonitoringServiceLevelObjective(request)

    @classmethod
    def list(self, project, service, service_account_file=""):
        stub = (
            service_level_objective_pb2_grpc.MonitoringServiceLevelObjectiveServiceStub(
                channel.Channel()
            )
        )
        request = (
            service_level_objective_pb2.ListMonitoringServiceLevelObjectiveRequest()
        )
        request.service_account_file = service_account_file
        request.Project = project

        request.Service = service

        return stub.ListMonitoringServiceLevelObjective(request).items

    def to_proto(self):
        resource = service_level_objective_pb2.MonitoringServiceLevelObjective()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.display_name):
            resource.display_name = Primitive.to_proto(self.display_name)
        if ServiceLevelObjectiveServiceLevelIndicator.to_proto(
            self.service_level_indicator
        ):
            resource.service_level_indicator.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicator.to_proto(
                    self.service_level_indicator
                )
            )
        else:
            resource.ClearField("service_level_indicator")
        if Primitive.to_proto(self.goal):
            resource.goal = Primitive.to_proto(self.goal)
        if Primitive.to_proto(self.rolling_period):
            resource.rolling_period = Primitive.to_proto(self.rolling_period)
        if ServiceLevelObjectiveCalendarPeriodEnum.to_proto(self.calendar_period):
            resource.calendar_period = ServiceLevelObjectiveCalendarPeriodEnum.to_proto(
                self.calendar_period
            )
        if Primitive.to_proto(self.user_labels):
            resource.user_labels = Primitive.to_proto(self.user_labels)
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        if Primitive.to_proto(self.service):
            resource.service = Primitive.to_proto(self.service)
        return resource


class ServiceLevelObjectiveServiceLevelIndicator(object):
    def __init__(
        self,
        basic_sli: dict = None,
        request_based: dict = None,
        windows_based: dict = None,
    ):
        self.basic_sli = basic_sli
        self.request_based = request_based
        self.windows_based = windows_based

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicator()
        )
        if ServiceLevelObjectiveServiceLevelIndicatorBasicSli.to_proto(
            resource.basic_sli
        ):
            res.basic_sli.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorBasicSli.to_proto(
                    resource.basic_sli
                )
            )
        else:
            res.ClearField("basic_sli")
        if ServiceLevelObjectiveServiceLevelIndicatorRequestBased.to_proto(
            resource.request_based
        ):
            res.request_based.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorRequestBased.to_proto(
                    resource.request_based
                )
            )
        else:
            res.ClearField("request_based")
        if ServiceLevelObjectiveServiceLevelIndicatorWindowsBased.to_proto(
            resource.windows_based
        ):
            res.windows_based.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorWindowsBased.to_proto(
                    resource.windows_based
                )
            )
        else:
            res.ClearField("windows_based")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicator(
            basic_sli=ServiceLevelObjectiveServiceLevelIndicatorBasicSli.from_proto(
                resource.basic_sli
            ),
            request_based=ServiceLevelObjectiveServiceLevelIndicatorRequestBased.from_proto(
                resource.request_based
            ),
            windows_based=ServiceLevelObjectiveServiceLevelIndicatorWindowsBased.from_proto(
                resource.windows_based
            ),
        )


class ServiceLevelObjectiveServiceLevelIndicatorArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicator.to_proto(i) for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicator.from_proto(i) for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorBasicSli(object):
    def __init__(
        self,
        method: list = None,
        location: list = None,
        version: list = None,
        availability: dict = None,
        latency: dict = None,
        operation_availability: dict = None,
        operation_latency: dict = None,
    ):
        self.method = method
        self.location = location
        self.version = version
        self.availability = availability
        self.latency = latency
        self.operation_availability = operation_availability
        self.operation_latency = operation_latency

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSli()
        )
        if Primitive.to_proto(resource.method):
            res.method.extend(Primitive.to_proto(resource.method))
        if Primitive.to_proto(resource.location):
            res.location.extend(Primitive.to_proto(resource.location))
        if Primitive.to_proto(resource.version):
            res.version.extend(Primitive.to_proto(resource.version))
        if ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability.to_proto(
            resource.availability
        ):
            res.availability.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability.to_proto(
                    resource.availability
                )
            )
        else:
            res.ClearField("availability")
        if ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency.to_proto(
            resource.latency
        ):
            res.latency.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency.to_proto(
                    resource.latency
                )
            )
        else:
            res.ClearField("latency")
        if ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability.to_proto(
            resource.operation_availability
        ):
            res.operation_availability.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability.to_proto(
                    resource.operation_availability
                )
            )
        else:
            res.ClearField("operation_availability")
        if ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency.to_proto(
            resource.operation_latency
        ):
            res.operation_latency.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency.to_proto(
                    resource.operation_latency
                )
            )
        else:
            res.ClearField("operation_latency")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorBasicSli(
            method=Primitive.from_proto(resource.method),
            location=Primitive.from_proto(resource.location),
            version=Primitive.from_proto(resource.version),
            availability=ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability.from_proto(
                resource.availability
            ),
            latency=ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency.from_proto(
                resource.latency
            ),
            operation_availability=ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability.from_proto(
                resource.operation_availability
            ),
            operation_latency=ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency.from_proto(
                resource.operation_latency
            ),
        )


class ServiceLevelObjectiveServiceLevelIndicatorBasicSliArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorBasicSli.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorBasicSli.from_proto(i)
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability()


class ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailabilityArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorBasicSliAvailability.from_proto(i)
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency(object):
    def __init__(self, threshold: str = None, experience: str = None):
        self.threshold = threshold
        self.experience = experience

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency()
        )
        if Primitive.to_proto(resource.threshold):
            res.threshold = Primitive.to_proto(resource.threshold)
        if ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum.to_proto(
            resource.experience
        ):
            res.experience = ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum.to_proto(
                resource.experience
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency(
            threshold=Primitive.from_proto(resource.threshold),
            experience=ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum.from_proto(
                resource.experience
            ),
        )


class ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatency.from_proto(i)
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability()


class ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailabilityArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationAvailability.from_proto(
                i
            )
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency(object):
    def __init__(self, threshold: str = None, experience: str = None):
        self.threshold = threshold
        self.experience = experience

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency()
        )
        if Primitive.to_proto(resource.threshold):
            res.threshold = Primitive.to_proto(resource.threshold)
        if ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum.to_proto(
            resource.experience
        ):
            res.experience = ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum.to_proto(
                resource.experience
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency(
            threshold=Primitive.from_proto(resource.threshold),
            experience=ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum.from_proto(
                resource.experience
            ),
        )


class ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatency.from_proto(
                i
            )
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorRequestBased(object):
    def __init__(self, good_total_ratio: dict = None, distribution_cut: dict = None):
        self.good_total_ratio = good_total_ratio
        self.distribution_cut = distribution_cut

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBased()
        )
        if ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio.to_proto(
            resource.good_total_ratio
        ):
            res.good_total_ratio.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio.to_proto(
                    resource.good_total_ratio
                )
            )
        else:
            res.ClearField("good_total_ratio")
        if ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut.to_proto(
            resource.distribution_cut
        ):
            res.distribution_cut.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut.to_proto(
                    resource.distribution_cut
                )
            )
        else:
            res.ClearField("distribution_cut")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorRequestBased(
            good_total_ratio=ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio.from_proto(
                resource.good_total_ratio
            ),
            distribution_cut=ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut.from_proto(
                resource.distribution_cut
            ),
        )


class ServiceLevelObjectiveServiceLevelIndicatorRequestBasedArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorRequestBased.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorRequestBased.from_proto(i)
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio(object):
    def __init__(
        self,
        good_service_filter: str = None,
        bad_service_filter: str = None,
        total_service_filter: str = None,
    ):
        self.good_service_filter = good_service_filter
        self.bad_service_filter = bad_service_filter
        self.total_service_filter = total_service_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio()
        )
        if Primitive.to_proto(resource.good_service_filter):
            res.good_service_filter = Primitive.to_proto(resource.good_service_filter)
        if Primitive.to_proto(resource.bad_service_filter):
            res.bad_service_filter = Primitive.to_proto(resource.bad_service_filter)
        if Primitive.to_proto(resource.total_service_filter):
            res.total_service_filter = Primitive.to_proto(resource.total_service_filter)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio(
            good_service_filter=Primitive.from_proto(resource.good_service_filter),
            bad_service_filter=Primitive.from_proto(resource.bad_service_filter),
            total_service_filter=Primitive.from_proto(resource.total_service_filter),
        )


class ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatioArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorRequestBasedGoodTotalRatio.from_proto(
                i
            )
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut(object):
    def __init__(self, distribution_filter: str = None, range: dict = None):
        self.distribution_filter = distribution_filter
        self.range = range

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut()
        )
        if Primitive.to_proto(resource.distribution_filter):
            res.distribution_filter = Primitive.to_proto(resource.distribution_filter)
        if ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange.to_proto(
            resource.range
        ):
            res.range.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange.to_proto(
                    resource.range
                )
            )
        else:
            res.ClearField("range")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut(
            distribution_filter=Primitive.from_proto(resource.distribution_filter),
            range=ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange.from_proto(
                resource.range
            ),
        )


class ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCut.from_proto(
                i
            )
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange(
    object
):
    def __init__(self, min: float = None, max: float = None):
        self.min = min
        self.max = max

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange()
        )
        if Primitive.to_proto(resource.min):
            res.min = Primitive.to_proto(resource.min)
        if Primitive.to_proto(resource.max):
            res.max = Primitive.to_proto(resource.max)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return (
            ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange(
                min=Primitive.from_proto(resource.min),
                max=Primitive.from_proto(resource.max),
            )
        )


class ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRangeArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorRequestBasedDistributionCutRange.from_proto(
                i
            )
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBased(object):
    def __init__(
        self,
        good_bad_metric_filter: str = None,
        good_total_ratio_threshold: dict = None,
        metric_mean_in_range: dict = None,
        metric_sum_in_range: dict = None,
        window_period: str = None,
    ):
        self.good_bad_metric_filter = good_bad_metric_filter
        self.good_total_ratio_threshold = good_total_ratio_threshold
        self.metric_mean_in_range = metric_mean_in_range
        self.metric_sum_in_range = metric_sum_in_range
        self.window_period = window_period

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBased()
        )
        if Primitive.to_proto(resource.good_bad_metric_filter):
            res.good_bad_metric_filter = Primitive.to_proto(
                resource.good_bad_metric_filter
            )
        if ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold.to_proto(
            resource.good_total_ratio_threshold
        ):
            res.good_total_ratio_threshold.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold.to_proto(
                    resource.good_total_ratio_threshold
                )
            )
        else:
            res.ClearField("good_total_ratio_threshold")
        if ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange.to_proto(
            resource.metric_mean_in_range
        ):
            res.metric_mean_in_range.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange.to_proto(
                    resource.metric_mean_in_range
                )
            )
        else:
            res.ClearField("metric_mean_in_range")
        if ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange.to_proto(
            resource.metric_sum_in_range
        ):
            res.metric_sum_in_range.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange.to_proto(
                    resource.metric_sum_in_range
                )
            )
        else:
            res.ClearField("metric_sum_in_range")
        if Primitive.to_proto(resource.window_period):
            res.window_period = Primitive.to_proto(resource.window_period)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorWindowsBased(
            good_bad_metric_filter=Primitive.from_proto(
                resource.good_bad_metric_filter
            ),
            good_total_ratio_threshold=ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold.from_proto(
                resource.good_total_ratio_threshold
            ),
            metric_mean_in_range=ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange.from_proto(
                resource.metric_mean_in_range
            ),
            metric_sum_in_range=ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange.from_proto(
                resource.metric_sum_in_range
            ),
            window_period=Primitive.from_proto(resource.window_period),
        )


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBased.to_proto(i)
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBased.from_proto(i)
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold(
    object
):
    def __init__(
        self,
        performance: dict = None,
        basic_sli_performance: dict = None,
        threshold: float = None,
    ):
        self.performance = performance
        self.basic_sli_performance = basic_sli_performance
        self.threshold = threshold

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold()
        )
        if ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance.to_proto(
            resource.performance
        ):
            res.performance.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance.to_proto(
                    resource.performance
                )
            )
        else:
            res.ClearField("performance")
        if ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance.to_proto(
            resource.basic_sli_performance
        ):
            res.basic_sli_performance.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance.to_proto(
                    resource.basic_sli_performance
                )
            )
        else:
            res.ClearField("basic_sli_performance")
        if Primitive.to_proto(resource.threshold):
            res.threshold = Primitive.to_proto(resource.threshold)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold(
            performance=ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance.from_proto(
                resource.performance
            ),
            basic_sli_performance=ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance.from_proto(
                resource.basic_sli_performance
            ),
            threshold=Primitive.from_proto(resource.threshold),
        )


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThreshold.from_proto(
                i
            )
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance(
    object
):
    def __init__(self, good_total_ratio: dict = None, distribution_cut: dict = None):
        self.good_total_ratio = good_total_ratio
        self.distribution_cut = distribution_cut

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance()
        )
        if ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio.to_proto(
            resource.good_total_ratio
        ):
            res.good_total_ratio.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio.to_proto(
                    resource.good_total_ratio
                )
            )
        else:
            res.ClearField("good_total_ratio")
        if ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut.to_proto(
            resource.distribution_cut
        ):
            res.distribution_cut.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut.to_proto(
                    resource.distribution_cut
                )
            )
        else:
            res.ClearField("distribution_cut")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance(
            good_total_ratio=ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio.from_proto(
                resource.good_total_ratio
            ),
            distribution_cut=ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut.from_proto(
                resource.distribution_cut
            ),
        )


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformance.from_proto(
                i
            )
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio(
    object
):
    def __init__(
        self,
        good_service_filter: str = None,
        bad_service_filter: str = None,
        total_service_filter: str = None,
    ):
        self.good_service_filter = good_service_filter
        self.bad_service_filter = bad_service_filter
        self.total_service_filter = total_service_filter

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio()
        )
        if Primitive.to_proto(resource.good_service_filter):
            res.good_service_filter = Primitive.to_proto(resource.good_service_filter)
        if Primitive.to_proto(resource.bad_service_filter):
            res.bad_service_filter = Primitive.to_proto(resource.bad_service_filter)
        if Primitive.to_proto(resource.total_service_filter):
            res.total_service_filter = Primitive.to_proto(resource.total_service_filter)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio(
            good_service_filter=Primitive.from_proto(resource.good_service_filter),
            bad_service_filter=Primitive.from_proto(resource.bad_service_filter),
            total_service_filter=Primitive.from_proto(resource.total_service_filter),
        )


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatioArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceGoodTotalRatio.from_proto(
                i
            )
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut(
    object
):
    def __init__(self, distribution_filter: str = None, range: dict = None):
        self.distribution_filter = distribution_filter
        self.range = range

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut()
        )
        if Primitive.to_proto(resource.distribution_filter):
            res.distribution_filter = Primitive.to_proto(resource.distribution_filter)
        if ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange.to_proto(
            resource.range
        ):
            res.range.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange.to_proto(
                    resource.range
                )
            )
        else:
            res.ClearField("range")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut(
            distribution_filter=Primitive.from_proto(resource.distribution_filter),
            range=ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange.from_proto(
                resource.range
            ),
        )


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCut.from_proto(
                i
            )
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange(
    object
):
    def __init__(self, min: float = None, max: float = None):
        self.min = min
        self.max = max

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange()
        )
        if Primitive.to_proto(resource.min):
            res.min = Primitive.to_proto(resource.min)
        if Primitive.to_proto(resource.max):
            res.max = Primitive.to_proto(resource.max)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange(
            min=Primitive.from_proto(resource.min),
            max=Primitive.from_proto(resource.max),
        )


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRangeArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdPerformanceDistributionCutRange.from_proto(
                i
            )
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance(
    object
):
    def __init__(
        self,
        method: list = None,
        location: list = None,
        version: list = None,
        availability: dict = None,
        latency: dict = None,
        operation_availability: dict = None,
        operation_latency: dict = None,
    ):
        self.method = method
        self.location = location
        self.version = version
        self.availability = availability
        self.latency = latency
        self.operation_availability = operation_availability
        self.operation_latency = operation_latency

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance()
        )
        if Primitive.to_proto(resource.method):
            res.method.extend(Primitive.to_proto(resource.method))
        if Primitive.to_proto(resource.location):
            res.location.extend(Primitive.to_proto(resource.location))
        if Primitive.to_proto(resource.version):
            res.version.extend(Primitive.to_proto(resource.version))
        if ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability.to_proto(
            resource.availability
        ):
            res.availability.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability.to_proto(
                    resource.availability
                )
            )
        else:
            res.ClearField("availability")
        if ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency.to_proto(
            resource.latency
        ):
            res.latency.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency.to_proto(
                    resource.latency
                )
            )
        else:
            res.ClearField("latency")
        if ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability.to_proto(
            resource.operation_availability
        ):
            res.operation_availability.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability.to_proto(
                    resource.operation_availability
                )
            )
        else:
            res.ClearField("operation_availability")
        if ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency.to_proto(
            resource.operation_latency
        ):
            res.operation_latency.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency.to_proto(
                    resource.operation_latency
                )
            )
        else:
            res.ClearField("operation_latency")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance(
            method=Primitive.from_proto(resource.method),
            location=Primitive.from_proto(resource.location),
            version=Primitive.from_proto(resource.version),
            availability=ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability.from_proto(
                resource.availability
            ),
            latency=ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency.from_proto(
                resource.latency
            ),
            operation_availability=ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability.from_proto(
                resource.operation_availability
            ),
            operation_latency=ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency.from_proto(
                resource.operation_latency
            ),
        )


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformance.from_proto(
                i
            )
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return (
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability()
        )


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailabilityArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceAvailability.from_proto(
                i
            )
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency(
    object
):
    def __init__(self, threshold: str = None, experience: str = None):
        self.threshold = threshold
        self.experience = experience

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency()
        )
        if Primitive.to_proto(resource.threshold):
            res.threshold = Primitive.to_proto(resource.threshold)
        if ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum.to_proto(
            resource.experience
        ):
            res.experience = ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum.to_proto(
                resource.experience
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency(
            threshold=Primitive.from_proto(resource.threshold),
            experience=ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum.from_proto(
                resource.experience
            ),
        )


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatency.from_proto(
                i
            )
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability()
        )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return (
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability()
        )


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailabilityArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationAvailability.from_proto(
                i
            )
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency(
    object
):
    def __init__(self, threshold: str = None, experience: str = None):
        self.threshold = threshold
        self.experience = experience

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency()
        )
        if Primitive.to_proto(resource.threshold):
            res.threshold = Primitive.to_proto(resource.threshold)
        if ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum.to_proto(
            resource.experience
        ):
            res.experience = ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum.to_proto(
                resource.experience
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency(
            threshold=Primitive.from_proto(resource.threshold),
            experience=ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum.from_proto(
                resource.experience
            ),
        )


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatency.from_proto(
                i
            )
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange(object):
    def __init__(self, time_series: str = None, range: dict = None):
        self.time_series = time_series
        self.range = range

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange()
        )
        if Primitive.to_proto(resource.time_series):
            res.time_series = Primitive.to_proto(resource.time_series)
        if ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange.to_proto(
            resource.range
        ):
            res.range.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange.to_proto(
                    resource.range
                )
            )
        else:
            res.ClearField("range")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange(
            time_series=Primitive.from_proto(resource.time_series),
            range=ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange.from_proto(
                resource.range
            ),
        )


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRange.from_proto(
                i
            )
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange(
    object
):
    def __init__(self, min: float = None, max: float = None):
        self.min = min
        self.max = max

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange()
        )
        if Primitive.to_proto(resource.min):
            res.min = Primitive.to_proto(resource.min)
        if Primitive.to_proto(resource.max):
            res.max = Primitive.to_proto(resource.max)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange(
            min=Primitive.from_proto(resource.min),
            max=Primitive.from_proto(resource.max),
        )


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRangeArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricMeanInRangeRange.from_proto(
                i
            )
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange(object):
    def __init__(self, time_series: str = None, range: dict = None):
        self.time_series = time_series
        self.range = range

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange()
        )
        if Primitive.to_proto(resource.time_series):
            res.time_series = Primitive.to_proto(resource.time_series)
        if ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange.to_proto(
            resource.range
        ):
            res.range.CopyFrom(
                ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange.to_proto(
                    resource.range
                )
            )
        else:
            res.ClearField("range")
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange(
            time_series=Primitive.from_proto(resource.time_series),
            range=ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange.from_proto(
                resource.range
            ),
        )


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRange.from_proto(
                i
            )
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange(
    object
):
    def __init__(self, min: float = None, max: float = None):
        self.min = min
        self.max = max

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = (
            service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange()
        )
        if Primitive.to_proto(resource.min):
            res.min = Primitive.to_proto(resource.min)
        if Primitive.to_proto(resource.max):
            res.max = Primitive.to_proto(resource.max)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return (
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange(
                min=Primitive.from_proto(resource.min),
                max=Primitive.from_proto(resource.max),
            )
        )


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRangeArray(
    object
):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange.to_proto(
                i
            )
            for i in resources
        ]

    @classmethod
    def from_proto(self, resources):
        return [
            ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedMetricSumInRangeRange.from_proto(
                i
            )
            for i in resources
        ]


class ServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum.Value(
            "MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum.Name(
            resource
        )[
            len(
                "MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliLatencyExperienceEnum"
            ) :
        ]


class ServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum.Value(
            "MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum.Name(
            resource
        )[
            len(
                "MonitoringServiceLevelObjectiveServiceLevelIndicatorBasicSliOperationLatencyExperienceEnum"
            ) :
        ]


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum.Value(
            "MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum.Name(
            resource
        )[
            len(
                "MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceLatencyExperienceEnum"
            ) :
        ]


class ServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum(
    object
):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum.Value(
            "MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum%s"
            % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_level_objective_pb2.MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum.Name(
            resource
        )[
            len(
                "MonitoringServiceLevelObjectiveServiceLevelIndicatorWindowsBasedGoodTotalRatioThresholdBasicSliPerformanceOperationLatencyExperienceEnum"
            ) :
        ]


class ServiceLevelObjectiveCalendarPeriodEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return service_level_objective_pb2.MonitoringServiceLevelObjectiveCalendarPeriodEnum.Value(
            "MonitoringServiceLevelObjectiveCalendarPeriodEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return service_level_objective_pb2.MonitoringServiceLevelObjectiveCalendarPeriodEnum.Name(
            resource
        )[
            len("MonitoringServiceLevelObjectiveCalendarPeriodEnum") :
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
