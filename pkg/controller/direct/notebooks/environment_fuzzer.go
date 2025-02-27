// Copyright 2024 Google LLC
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
// proto.message: google.cloud.notebooks.v1.Environment
// api.group: notebooks.cnrm.cloud.google.com

package notebooks

import (
	pb "cloud.google.com/go/notebooks/apiv1beta1/notebookspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(notebooksEnvironmentFuzzer())
}

func notebooksEnvironmentFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Environment{},
		NotebooksEnvironmentSpec_FromProto, NotebooksEnvironmentSpec_ToProto,
		NotebooksEnvironmentObservedState_FromProto, NotebooksEnvironmentObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".vm_image")
	f.SpecFields.Insert(".container_image")
	f.SpecFields.Insert(".post_startup_script")

	f.StatusFields.Insert(".create_time")

	return f
}
