// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.logging.v2.LogMetric
// api.group: logging.cnrm.cloud.google.com

package logging

import (
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(loggingLogMetricFuzzer())
}

func loggingLogMetricFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.LogMetric{},
		LoggingLogMetricSpec_FromProto, LoggingLogMetricSpec_ToProto,
		LoggingLogMetricStatus_FromProto, LoggingLogMetricStatus_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".filter")
	f.SpecField(".disabled")
	f.SpecField(".value_extractor")
	f.SpecField(".label_extractors")

	f.SpecField(".metric_descriptor.display_name")
	f.SpecField(".metric_descriptor.labels")
	f.SpecField(".metric_descriptor.launch_stage")
	f.SpecField(".metric_descriptor.metric_kind")
	f.SpecField(".metric_descriptor.unit")
	f.SpecField(".metric_descriptor.value_type")

	f.SpecField(".bucket_options")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".metric_descriptor.description")
	f.StatusField(".metric_descriptor.monitored_resource_types")
	f.StatusField(".metric_descriptor.name")
	f.StatusField(".metric_descriptor.type")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_Internal(".bucket_name")
	f.Unimplemented_Internal(".version")
	f.Unimplemented_NotYetTriaged(".metric_descriptor.metadata")

	return f
}
