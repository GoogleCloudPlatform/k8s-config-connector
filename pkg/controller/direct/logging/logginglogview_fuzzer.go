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
// proto.message: google.logging.v2.LogView
// api.group: logging.cnrm.cloud.google.com

package logging

import (
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(loggingLogViewFuzzer())
}

func loggingLogViewFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.LogView{},
		LoggingLogViewSpec_FromProto, LoggingLogViewSpec_ToProto,
		LoggingLogViewStatus_FromProto, LoggingLogViewStatus_ToProto,
	)

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".filter")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")

	f.UnimplementedFields.Insert(".name")

	return f
}
