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
// proto.message: google.cloud.developerconnect.v1.InsightsConfig
// api.group: developerconnect.cnrm.cloud.google.com

package developerconnect

import (
	pb "cloud.google.com/go/developerconnect/apiv1/developerconnectpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(devconnectInsightsConfigFuzzer())
}

func devconnectInsightsConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.InsightsConfig{},
		DevConnectInsightsConfigSpec_FromProto, DevConnectInsightsConfigSpec_ToProto,
		DevConnectInsightsConfigObservedState_FromProto, DevConnectInsightsConfigObservedState_ToProto,
	)

	f.SpecField(".app_hub_application")
	f.SpecField(".artifact_configs")

	f.Unimplemented_NotYetTriaged(".annotations")
	f.Unimplemented_NotYetTriaged(".labels")
	f.Unimplemented_NotYetTriaged(".projects")
	f.Unimplemented_NotYetTriaged(".errors[].details")
	f.Unimplemented_NotYetTriaged(".runtime_configs[].app_hub_service")
	f.Unimplemented_NotYetTriaged(".runtime_configs[].google_cloud_run")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".runtime_configs")
	f.StatusField(".state")
	f.StatusField(".reconciling")
	f.StatusField(".errors")

	f.Unimplemented_Identity(".name")

	return f
}
