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
// proto.message: google.logging.v2.LogBucket
// api.group: logging.cnrm.cloud.google.com

package logging

import (
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(loggingLogBucketFuzzer())
}

func loggingLogBucketFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.LogBucket{},
		LoggingLogBucketSpec_FromProto, LoggingLogBucketSpec_ToProto,
		LoggingLogBucketStatus_FromProto, LoggingLogBucketStatus_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".retention_days")
	f.SpecField(".locked")
	f.SpecField(".analytics_enabled")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".lifecycle_state")

	f.Unimplemented_Identity(".name")

	f.Unimplemented_NotYetTriaged(".restricted_fields")
	f.Unimplemented_NotYetTriaged(".cmek_settings")
	f.Unimplemented_NotYetTriaged(".index_configs")

	return f
}
