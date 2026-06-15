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
// proto.message: google.logging.v2.LogSink
// api.group: logging.cnrm.cloud.google.com

package logging

import (
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(loggingLogSinkFuzzer())
}

func loggingLogSinkFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.LogSink{},
		LoggingLogSinkSpec_FromProto, LoggingLogSinkSpec_ToProto,
		LoggingLogSinkStatus_FromProto, LoggingLogSinkStatus_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".destination")
	f.SpecField(".filter")
	f.SpecField(".disabled")
	f.SpecField(".exclusions")
	f.SpecField(".include_children")
	f.SpecField(".bigquery_options")

	f.StatusField(".writer_identity")

	f.Unimplemented_Identity(".name")
	f.Unimplemented_NotYetTriaged(".output_version_format")

	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".update_time")
	f.Unimplemented_NotYetTriaged(".exclusions[].create_time")
	f.Unimplemented_NotYetTriaged(".exclusions[].update_time")
	f.Unimplemented_NotYetTriaged(".bigquery_options.uses_timestamp_column_partitioning")

	f.FilterSpec = func(in *pb.LogSink) {
		if in.Destination != "" {
			switch in.Destination[0] % 4 {
			case 0:
				in.Destination = "bigquery.googleapis.com/" + in.Destination
			case 1:
				in.Destination = "logging.googleapis.com/" + in.Destination
			case 2:
				in.Destination = "pubsub.googleapis.com/" + in.Destination
			default:
				in.Destination = "storage.googleapis.com/" + in.Destination
			}
		}
	}

	return f
}
