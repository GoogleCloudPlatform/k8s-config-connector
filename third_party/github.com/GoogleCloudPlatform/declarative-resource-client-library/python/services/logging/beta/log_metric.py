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
from google3.cloud.graphite.mmv2.services.google.logging import log_metric_pb2
from google3.cloud.graphite.mmv2.services.google.logging import log_metric_pb2_grpc

from typing import List


class LogMetric(object):
    def __init__(
        self,
        name: str = None,
        description: str = None,
        filter: str = None,
        disabled: bool = None,
        metric_descriptor: dict = None,
        value_extractor: str = None,
        label_extractors: dict = None,
        bucket_options: dict = None,
        create_time: str = None,
        update_time: str = None,
        project: str = None,
        service_account_file: str = "",
    ):
        channel.initialize()
        self.name = name
        self.description = description
        self.filter = filter
        self.disabled = disabled
        self.metric_descriptor = metric_descriptor
        self.value_extractor = value_extractor
        self.label_extractors = label_extractors
        self.bucket_options = bucket_options
        self.project = project
        self.service_account_file = service_account_file

    def apply(self):
        stub = log_metric_pb2_grpc.LoggingBetaLogMetricServiceStub(channel.Channel())
        request = log_metric_pb2.ApplyLoggingBetaLogMetricRequest()
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.filter):
            request.resource.filter = Primitive.to_proto(self.filter)

        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if LogMetricMetricDescriptor.to_proto(self.metric_descriptor):
            request.resource.metric_descriptor.CopyFrom(
                LogMetricMetricDescriptor.to_proto(self.metric_descriptor)
            )
        else:
            request.resource.ClearField("metric_descriptor")
        if Primitive.to_proto(self.value_extractor):
            request.resource.value_extractor = Primitive.to_proto(self.value_extractor)

        if Primitive.to_proto(self.label_extractors):
            request.resource.label_extractors = Primitive.to_proto(
                self.label_extractors
            )

        if LogMetricBucketOptions.to_proto(self.bucket_options):
            request.resource.bucket_options.CopyFrom(
                LogMetricBucketOptions.to_proto(self.bucket_options)
            )
        else:
            request.resource.ClearField("bucket_options")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        request.service_account_file = self.service_account_file

        response = stub.ApplyLoggingBetaLogMetric(request)
        self.name = Primitive.from_proto(response.name)
        self.description = Primitive.from_proto(response.description)
        self.filter = Primitive.from_proto(response.filter)
        self.disabled = Primitive.from_proto(response.disabled)
        self.metric_descriptor = LogMetricMetricDescriptor.from_proto(
            response.metric_descriptor
        )
        self.value_extractor = Primitive.from_proto(response.value_extractor)
        self.label_extractors = Primitive.from_proto(response.label_extractors)
        self.bucket_options = LogMetricBucketOptions.from_proto(response.bucket_options)
        self.create_time = Primitive.from_proto(response.create_time)
        self.update_time = Primitive.from_proto(response.update_time)
        self.project = Primitive.from_proto(response.project)

    def delete(self):
        stub = log_metric_pb2_grpc.LoggingBetaLogMetricServiceStub(channel.Channel())
        request = log_metric_pb2.DeleteLoggingBetaLogMetricRequest()
        request.service_account_file = self.service_account_file
        if Primitive.to_proto(self.name):
            request.resource.name = Primitive.to_proto(self.name)

        if Primitive.to_proto(self.description):
            request.resource.description = Primitive.to_proto(self.description)

        if Primitive.to_proto(self.filter):
            request.resource.filter = Primitive.to_proto(self.filter)

        if Primitive.to_proto(self.disabled):
            request.resource.disabled = Primitive.to_proto(self.disabled)

        if LogMetricMetricDescriptor.to_proto(self.metric_descriptor):
            request.resource.metric_descriptor.CopyFrom(
                LogMetricMetricDescriptor.to_proto(self.metric_descriptor)
            )
        else:
            request.resource.ClearField("metric_descriptor")
        if Primitive.to_proto(self.value_extractor):
            request.resource.value_extractor = Primitive.to_proto(self.value_extractor)

        if Primitive.to_proto(self.label_extractors):
            request.resource.label_extractors = Primitive.to_proto(
                self.label_extractors
            )

        if LogMetricBucketOptions.to_proto(self.bucket_options):
            request.resource.bucket_options.CopyFrom(
                LogMetricBucketOptions.to_proto(self.bucket_options)
            )
        else:
            request.resource.ClearField("bucket_options")
        if Primitive.to_proto(self.project):
            request.resource.project = Primitive.to_proto(self.project)

        response = stub.DeleteLoggingBetaLogMetric(request)

    @classmethod
    def list(self, project, service_account_file=""):
        stub = log_metric_pb2_grpc.LoggingBetaLogMetricServiceStub(channel.Channel())
        request = log_metric_pb2.ListLoggingBetaLogMetricRequest()
        request.service_account_file = service_account_file
        request.Project = project

        return stub.ListLoggingBetaLogMetric(request).items

    def to_proto(self):
        resource = log_metric_pb2.LoggingBetaLogMetric()
        if Primitive.to_proto(self.name):
            resource.name = Primitive.to_proto(self.name)
        if Primitive.to_proto(self.description):
            resource.description = Primitive.to_proto(self.description)
        if Primitive.to_proto(self.filter):
            resource.filter = Primitive.to_proto(self.filter)
        if Primitive.to_proto(self.disabled):
            resource.disabled = Primitive.to_proto(self.disabled)
        if LogMetricMetricDescriptor.to_proto(self.metric_descriptor):
            resource.metric_descriptor.CopyFrom(
                LogMetricMetricDescriptor.to_proto(self.metric_descriptor)
            )
        else:
            resource.ClearField("metric_descriptor")
        if Primitive.to_proto(self.value_extractor):
            resource.value_extractor = Primitive.to_proto(self.value_extractor)
        if Primitive.to_proto(self.label_extractors):
            resource.label_extractors = Primitive.to_proto(self.label_extractors)
        if LogMetricBucketOptions.to_proto(self.bucket_options):
            resource.bucket_options.CopyFrom(
                LogMetricBucketOptions.to_proto(self.bucket_options)
            )
        else:
            resource.ClearField("bucket_options")
        if Primitive.to_proto(self.project):
            resource.project = Primitive.to_proto(self.project)
        return resource


class LogMetricMetricDescriptor(object):
    def __init__(
        self,
        name: str = None,
        type: str = None,
        labels: list = None,
        metric_kind: str = None,
        value_type: str = None,
        unit: str = None,
        description: str = None,
        display_name: str = None,
        metadata: dict = None,
        launch_stage: str = None,
        monitored_resource_types: list = None,
    ):
        self.name = name
        self.type = type
        self.labels = labels
        self.metric_kind = metric_kind
        self.value_type = value_type
        self.unit = unit
        self.description = description
        self.display_name = display_name
        self.metadata = metadata
        self.launch_stage = launch_stage
        self.monitored_resource_types = monitored_resource_types

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = log_metric_pb2.LoggingBetaLogMetricMetricDescriptor()
        if Primitive.to_proto(resource.name):
            res.name = Primitive.to_proto(resource.name)
        if Primitive.to_proto(resource.type):
            res.type = Primitive.to_proto(resource.type)
        if LogMetricMetricDescriptorLabelsArray.to_proto(resource.labels):
            res.labels.extend(
                LogMetricMetricDescriptorLabelsArray.to_proto(resource.labels)
            )
        if LogMetricMetricDescriptorMetricKindEnum.to_proto(resource.metric_kind):
            res.metric_kind = LogMetricMetricDescriptorMetricKindEnum.to_proto(
                resource.metric_kind
            )
        if LogMetricMetricDescriptorValueTypeEnum.to_proto(resource.value_type):
            res.value_type = LogMetricMetricDescriptorValueTypeEnum.to_proto(
                resource.value_type
            )
        if Primitive.to_proto(resource.unit):
            res.unit = Primitive.to_proto(resource.unit)
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        if Primitive.to_proto(resource.display_name):
            res.display_name = Primitive.to_proto(resource.display_name)
        if LogMetricMetricDescriptorMetadata.to_proto(resource.metadata):
            res.metadata.CopyFrom(
                LogMetricMetricDescriptorMetadata.to_proto(resource.metadata)
            )
        else:
            res.ClearField("metadata")
        if LogMetricMetricDescriptorLaunchStageEnum.to_proto(resource.launch_stage):
            res.launch_stage = LogMetricMetricDescriptorLaunchStageEnum.to_proto(
                resource.launch_stage
            )
        if Primitive.to_proto(resource.monitored_resource_types):
            res.monitored_resource_types.extend(
                Primitive.to_proto(resource.monitored_resource_types)
            )
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return LogMetricMetricDescriptor(
            name=Primitive.from_proto(resource.name),
            type=Primitive.from_proto(resource.type),
            labels=LogMetricMetricDescriptorLabelsArray.from_proto(resource.labels),
            metric_kind=LogMetricMetricDescriptorMetricKindEnum.from_proto(
                resource.metric_kind
            ),
            value_type=LogMetricMetricDescriptorValueTypeEnum.from_proto(
                resource.value_type
            ),
            unit=Primitive.from_proto(resource.unit),
            description=Primitive.from_proto(resource.description),
            display_name=Primitive.from_proto(resource.display_name),
            metadata=LogMetricMetricDescriptorMetadata.from_proto(resource.metadata),
            launch_stage=LogMetricMetricDescriptorLaunchStageEnum.from_proto(
                resource.launch_stage
            ),
            monitored_resource_types=Primitive.from_proto(
                resource.monitored_resource_types
            ),
        )


class LogMetricMetricDescriptorArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [LogMetricMetricDescriptor.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [LogMetricMetricDescriptor.from_proto(i) for i in resources]


class LogMetricMetricDescriptorLabels(object):
    def __init__(
        self, key: str = None, value_type: str = None, description: str = None
    ):
        self.key = key
        self.value_type = value_type
        self.description = description

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = log_metric_pb2.LoggingBetaLogMetricMetricDescriptorLabels()
        if Primitive.to_proto(resource.key):
            res.key = Primitive.to_proto(resource.key)
        if LogMetricMetricDescriptorLabelsValueTypeEnum.to_proto(resource.value_type):
            res.value_type = LogMetricMetricDescriptorLabelsValueTypeEnum.to_proto(
                resource.value_type
            )
        if Primitive.to_proto(resource.description):
            res.description = Primitive.to_proto(resource.description)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return LogMetricMetricDescriptorLabels(
            key=Primitive.from_proto(resource.key),
            value_type=LogMetricMetricDescriptorLabelsValueTypeEnum.from_proto(
                resource.value_type
            ),
            description=Primitive.from_proto(resource.description),
        )


class LogMetricMetricDescriptorLabelsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [LogMetricMetricDescriptorLabels.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [LogMetricMetricDescriptorLabels.from_proto(i) for i in resources]


class LogMetricMetricDescriptorMetadata(object):
    def __init__(self, sample_period: str = None, ingest_delay: str = None):
        self.sample_period = sample_period
        self.ingest_delay = ingest_delay

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = log_metric_pb2.LoggingBetaLogMetricMetricDescriptorMetadata()
        if Primitive.to_proto(resource.sample_period):
            res.sample_period = Primitive.to_proto(resource.sample_period)
        if Primitive.to_proto(resource.ingest_delay):
            res.ingest_delay = Primitive.to_proto(resource.ingest_delay)
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return LogMetricMetricDescriptorMetadata(
            sample_period=Primitive.from_proto(resource.sample_period),
            ingest_delay=Primitive.from_proto(resource.ingest_delay),
        )


class LogMetricMetricDescriptorMetadataArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [LogMetricMetricDescriptorMetadata.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [LogMetricMetricDescriptorMetadata.from_proto(i) for i in resources]


class LogMetricBucketOptions(object):
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

        res = log_metric_pb2.LoggingBetaLogMetricBucketOptions()
        if LogMetricBucketOptionsLinearBuckets.to_proto(resource.linear_buckets):
            res.linear_buckets.CopyFrom(
                LogMetricBucketOptionsLinearBuckets.to_proto(resource.linear_buckets)
            )
        else:
            res.ClearField("linear_buckets")
        if LogMetricBucketOptionsExponentialBuckets.to_proto(
            resource.exponential_buckets
        ):
            res.exponential_buckets.CopyFrom(
                LogMetricBucketOptionsExponentialBuckets.to_proto(
                    resource.exponential_buckets
                )
            )
        else:
            res.ClearField("exponential_buckets")
        if LogMetricBucketOptionsExplicitBuckets.to_proto(resource.explicit_buckets):
            res.explicit_buckets.CopyFrom(
                LogMetricBucketOptionsExplicitBuckets.to_proto(
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

        return LogMetricBucketOptions(
            linear_buckets=LogMetricBucketOptionsLinearBuckets.from_proto(
                resource.linear_buckets
            ),
            exponential_buckets=LogMetricBucketOptionsExponentialBuckets.from_proto(
                resource.exponential_buckets
            ),
            explicit_buckets=LogMetricBucketOptionsExplicitBuckets.from_proto(
                resource.explicit_buckets
            ),
        )


class LogMetricBucketOptionsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [LogMetricBucketOptions.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [LogMetricBucketOptions.from_proto(i) for i in resources]


class LogMetricBucketOptionsLinearBuckets(object):
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

        res = log_metric_pb2.LoggingBetaLogMetricBucketOptionsLinearBuckets()
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

        return LogMetricBucketOptionsLinearBuckets(
            num_finite_buckets=Primitive.from_proto(resource.num_finite_buckets),
            width=Primitive.from_proto(resource.width),
            offset=Primitive.from_proto(resource.offset),
        )


class LogMetricBucketOptionsLinearBucketsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [LogMetricBucketOptionsLinearBuckets.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [LogMetricBucketOptionsLinearBuckets.from_proto(i) for i in resources]


class LogMetricBucketOptionsExponentialBuckets(object):
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

        res = log_metric_pb2.LoggingBetaLogMetricBucketOptionsExponentialBuckets()
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

        return LogMetricBucketOptionsExponentialBuckets(
            num_finite_buckets=Primitive.from_proto(resource.num_finite_buckets),
            growth_factor=Primitive.from_proto(resource.growth_factor),
            scale=Primitive.from_proto(resource.scale),
        )


class LogMetricBucketOptionsExponentialBucketsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [LogMetricBucketOptionsExponentialBuckets.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [
            LogMetricBucketOptionsExponentialBuckets.from_proto(i) for i in resources
        ]


class LogMetricBucketOptionsExplicitBuckets(object):
    def __init__(self, bounds: list = None):
        self.bounds = bounds

    @classmethod
    def to_proto(self, resource):
        if not resource:
            return None

        res = log_metric_pb2.LoggingBetaLogMetricBucketOptionsExplicitBuckets()
        if float64Array.to_proto(resource.bounds):
            res.bounds.extend(float64Array.to_proto(resource.bounds))
        return res

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return None

        return LogMetricBucketOptionsExplicitBuckets(
            bounds=float64Array.from_proto(resource.bounds),
        )


class LogMetricBucketOptionsExplicitBucketsArray(object):
    @classmethod
    def to_proto(self, resources):
        if not resources:
            return resources
        return [LogMetricBucketOptionsExplicitBuckets.to_proto(i) for i in resources]

    @classmethod
    def from_proto(self, resources):
        return [LogMetricBucketOptionsExplicitBuckets.from_proto(i) for i in resources]


class LogMetricMetricDescriptorLabelsValueTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return log_metric_pb2.LoggingBetaLogMetricMetricDescriptorLabelsValueTypeEnum.Value(
            "LoggingBetaLogMetricMetricDescriptorLabelsValueTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return (
            log_metric_pb2.LoggingBetaLogMetricMetricDescriptorLabelsValueTypeEnum.Name(
                resource
            )[len("LoggingBetaLogMetricMetricDescriptorLabelsValueTypeEnum") :]
        )


class LogMetricMetricDescriptorMetricKindEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return log_metric_pb2.LoggingBetaLogMetricMetricDescriptorMetricKindEnum.Value(
            "LoggingBetaLogMetricMetricDescriptorMetricKindEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return log_metric_pb2.LoggingBetaLogMetricMetricDescriptorMetricKindEnum.Name(
            resource
        )[len("LoggingBetaLogMetricMetricDescriptorMetricKindEnum") :]


class LogMetricMetricDescriptorValueTypeEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return log_metric_pb2.LoggingBetaLogMetricMetricDescriptorValueTypeEnum.Value(
            "LoggingBetaLogMetricMetricDescriptorValueTypeEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return log_metric_pb2.LoggingBetaLogMetricMetricDescriptorValueTypeEnum.Name(
            resource
        )[len("LoggingBetaLogMetricMetricDescriptorValueTypeEnum") :]


class LogMetricMetricDescriptorLaunchStageEnum(object):
    @classmethod
    def to_proto(self, resource):
        if not resource:
            return resource
        return log_metric_pb2.LoggingBetaLogMetricMetricDescriptorLaunchStageEnum.Value(
            "LoggingBetaLogMetricMetricDescriptorLaunchStageEnum%s" % resource
        )

    @classmethod
    def from_proto(self, resource):
        if not resource:
            return resource
        return log_metric_pb2.LoggingBetaLogMetricMetricDescriptorLaunchStageEnum.Name(
            resource
        )[len("LoggingBetaLogMetricMetricDescriptorLaunchStageEnum") :]


class Primitive(object):
    @classmethod
    def to_proto(self, s):
        if not s:
            return ""
        return s

    @classmethod
    def from_proto(self, s):
        return s
