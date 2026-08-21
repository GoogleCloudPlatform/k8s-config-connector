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
		NotebooksEnvironmentSpec_v1alpha1_FromProto, NotebooksEnvironmentSpec_v1alpha1_ToProto,
		NotebooksEnvironmentObservedState_v1alpha1_FromProto, NotebooksEnvironmentObservedState_v1alpha1_ToProto,
	)

	// Field comparison between KRM NotebooksEnvironment and Proto Environment:
	// - .spec.displayName       <=> .display_name
	// - .spec.description       <=> .description
	// - .spec.vmImage           <=> .vm_image
	// - .spec.containerImage    <=> .container_image
	// - .spec.postStartupScript <=> .post_startup_script
	//
	// Fields ignored because they are not part of KRM Spec:
	// - .name                   (handled by resource name/identity URI)
	// - .projectRef             (parent reference handled by KCC framework)
	// - .location               (region/zone handled by KCC framework)
	// - .resourceID             (handled by KCC framework)
	//
	// Status fields:
	// - .status.observedState.createTime <=> .create_time

	f.Unimplemented_Identity(".name")

	f.SpecField(".display_name")
	f.SpecField(".description")
	f.SpecField(".vm_image")
	f.SpecField(".container_image")
	f.SpecField(".post_startup_script")

	f.StatusField(".create_time")

	return f
}
