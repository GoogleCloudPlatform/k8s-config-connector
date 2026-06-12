// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.notebooks.v1.Runtime
// api.group: notebooks.cnrm.cloud.google.com

package notebooks

import (
	pb "cloud.google.com/go/notebooks/apiv1/notebookspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(notebookRuntimeFuzzer())
}

func notebookRuntimeFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Runtime{},
		Runtime_v1alpha1_FromProto, Runtime_v1alpha1_ToProto,
		RuntimeObservedState_v1alpha1_FromProto, RuntimeObservedState_v1alpha1_ToProto,
	)

	f.Unimplemented_Identity(".name")

	// Unimplemented fields because the KRM representation is minimal
	f.Unimplemented_NotYetTriaged(".state")
	f.Unimplemented_NotYetTriaged(".health_state")
	f.Unimplemented_NotYetTriaged(".access_config")
	f.Unimplemented_NotYetTriaged(".software_config")
	f.Unimplemented_NotYetTriaged(".metrics")
	f.Unimplemented_NotYetTriaged(".create_time")
	f.Unimplemented_NotYetTriaged(".update_time")
	f.Unimplemented_NotYetTriaged(".virtual_machine")

	return f
}
