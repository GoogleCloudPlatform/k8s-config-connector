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

package apihubcuration

import (
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/apihub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(apihubCurationFuzzer())
}

func apihubCurationFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Curation{},
		apihub.APIHubCurationSpec_FromProto, apihub.APIHubCurationSpec_ToProto,
		apihub.APIHubCurationObservedState_FromProto, apihub.APIHubCurationObservedState_ToProto,
	)

	// Identity Field
	f.Unimplemented_Identity(".name")

	// Spec fields
	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".endpoint")

	// Status fields
	f.StatusField(".plugin_instance_actions")
	f.StatusField(".last_execution_state")
	f.StatusField(".last_execution_error_code")
	f.StatusField(".last_execution_error_message")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	return f
}
