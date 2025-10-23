// Copyright 2025 Google LLC
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
// proto.message: google.cloud.deploy.v1.Target
// api.group: clouddeploy.googleapis.com

package clouddeploy

import (
	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(cloudDeployTargetFuzzer())
}

func cloudDeployTargetFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Target{},
		CloudDeployTargetSpec_FromProto, CloudDeployTargetSpec_ToProto,
		CloudDeployTargetObservedState_FromProto, CloudDeployTargetObservedState_ToProto,
	)

	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".annotations")
	f.SpecFields.Insert(".require_approval")
	f.SpecFields.Insert(".gke")
	f.SpecFields.Insert(".anthos_cluster")
	f.SpecFields.Insert(".run")
	f.SpecFields.Insert(".multi_target")
	f.SpecFields.Insert(".custom_target")
	f.SpecFields.Insert(".execution_configs")
	f.SpecFields.Insert(".deploy_parameters")
	f.SpecFields.Insert(".associated_entities")

	f.StatusFields.Insert(".target_id")
	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".etag")

	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".labels")

	return f
}
