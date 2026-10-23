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
// proto.message: google.logging.v2.LogScope
// api.group: logging.cnrm.cloud.google.com

package logging

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	api "google.golang.org/api/logging/v2"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(loggingLogScopeFuzzer())
}

func loggingLogScopeFuzzer() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&api.LogScope{},
		LoggingLogScopeSpec_FromAPI, LoggingLogScopeSpec_ToAPI,
		LoggingLogScopeStatus_FromAPI, LoggingLogScopeStatus_ToAPI,
	)

	f.SpecField(".Description")
	f.SpecField(".ResourceNames")

	f.StatusField(".CreateTime")
	f.StatusField(".UpdateTime")

	f.IdentityField(".Name")

	f.Ignore_JSONBookkeeping(".ForceSendFields")
	f.Ignore_JSONBookkeeping(".NullFields")
	f.Ignore_JSONBookkeeping(".ServerResponse")

	f.FilterSpec = func(in *api.LogScope) {
		for i, rn := range in.ResourceNames {
			in.ResourceNames[i] = "projects/" + rn
		}
	}

	return f
}
