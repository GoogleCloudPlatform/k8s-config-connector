// Copyright 2024 Google LLC
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
// api.group: deploy.cnrm.cloud.google.com

package deploy

import (
	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting/fuzzers/fieldtrimmer"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(deployTargetFuzzer())
}

func deployTargetFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Target{},
		DeployTargetSpec_FromProto, DeployTargetSpec_ToProto,
		DeployTargetObservedState_FromProto, DeployTargetObservedState_ToProto,
	)

	f.SpecFields.Insert(".gke")
	f.SpecFields.Insert(".anthos_cluster")
	f.SpecFields.Insert(".run")
	f.SpecFields.Insert(".multi_target")
	f.SpecFields.Insert(".custom_target")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".annotations")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".require_approval")
	f.SpecFields.Insert(".execution_configs")
	f.SpecFields.Insert(".deploy_parameters")

	f.StatusFields.Insert(".target_id")
	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")

	f.UnimplementedFields.Insert(".associated_entities") // Map of struct.
	f.UnimplementedFields.Insert(".etag")                // Computed checksum.
	f.UnimplementedFields.Insert(".name")                // Handled in config-connector code

	f.AddProtobufFuzzer(&fuzztesting.KstructFuzzer{}) // Field name is "gke".
	f.AddProtobufFuzzer(&fuzztesting.KstructFuzzer{}) // Field name is "run".

	// Trim fields with regex match.
	f.SpecFields.Add(fieldtrimmer.NewFieldTrimmer(
		".gke.cluster",
		`^projects/[^/]+/locations/[^/]+/clusters/[^/]+$`,
	))
	f.SpecFields.Add(fieldtrimmer.NewFieldTrimmer(
		".anthos_cluster.membership",
		`^projects/[^/]+/locations/[^/]+/memberships/[^/]+$`,
	))
	f.SpecFields.Add(fieldtrimmer.NewFieldTrimmer(
		".run.location",
		"^projects/[^/]+/locations/[^/]+$",
	))

	return f
}

```
</out>


