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
// proto.message: google.cloud.dataform.v1beta1.Folder
// api.group: dataform.cnrm.cloud.google.com

package dataform

import (
	dataformpb "cloud.google.com/go/dataform/apiv1beta1/dataformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(dataformFolderFuzzer())
}

func dataformFolderFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&dataformpb.Folder{},
		DataformFolderSpec_v1alpha1_FromProto, DataformFolderSpec_v1alpha1_ToProto,
		DataformFolderObservedState_v1alpha1_FromProto, DataformFolderObservedState_v1alpha1_ToProto,
	)

	// Identity fields that are not in KRM fields
	f.Unimplemented_Identity(".name")
	f.Unimplemented_Identity(".containing_folder")

	// Spec fields to fuzz
	f.SpecField(".display_name")

	// Status fields to fuzz (observed state)
	f.StatusField(".team_folder_name")
	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".internal_metadata")
	f.StatusField(".creator_iam_principal")

	return f
}
