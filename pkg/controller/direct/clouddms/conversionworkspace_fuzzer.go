// Copyright 2025 Google LLC
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

package clouddms

import (
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzCloudDMSConversionWorkspace())
}

func fuzzCloudDMSConversionWorkspace() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.ConversionWorkspace{},
		CloudDMSConversionWorkspaceSpec_FromProto, CloudDMSConversionWorkspaceSpec_ToProto,
		CloudDMSConversionWorkspaceObservedState_FromProto, CloudDMSConversionWorkspaceObservedState_ToProto,
	)
	f.UnimplementedFields.Insert(".name")

	f.SpecFields.Insert(".source")
	f.SpecFields.Insert(".source.engine")
	f.SpecFields.Insert(".source.version")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".destination")
	f.SpecFields.Insert(".destination.engine")
	f.SpecFields.Insert(".destination.version")
	f.SpecFields.Insert(".global_settings")
	f.SpecFields.Insert(".display_name")

	f.StatusFields.Insert(".has_uncommitted_changes")
	f.StatusFields.Insert(".latest_commit_id")
	f.StatusFields.Insert(".latest_commit_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".create_time")
	return f
}
