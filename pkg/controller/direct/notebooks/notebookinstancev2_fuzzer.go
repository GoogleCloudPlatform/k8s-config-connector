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
// proto.message: google.cloud.notebooks.v2.Instance
// api.group: notebooks.cnrm.cloud.google.com

package notebooks

import (
	pb "cloud.google.com/go/notebooks/apiv2/notebookspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(notebookInstanceV2Fuzzer())
}

func notebookInstanceV2Fuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Instance{},
		NotebookInstanceV2Spec_v1alpha1_FromProto, NotebookInstanceV2Spec_v1alpha1_ToProto,
		NotebookInstanceV2ObservedState_v1alpha1_FromProto, NotebookInstanceV2ObservedState_v1alpha1_ToProto,
	)

	f.Unimplemented_Identity(".name") // special field

	f.SpecField(".gce_setup")
	f.SpecField(".instance_owners")
	f.SpecField(".disable_proxy_access")
	f.SpecField(".labels")

	f.StatusField(".gce_setup")
	f.StatusField(".proxy_uri")
	f.StatusField(".creator")
	f.StatusField(".state")
	f.StatusField(".upgrade_history")
	f.StatusField(".health_state")
	f.StatusField(".health_info")
	f.StatusField(".create_time")
	f.StatusField(".update_time")

	return f
}
