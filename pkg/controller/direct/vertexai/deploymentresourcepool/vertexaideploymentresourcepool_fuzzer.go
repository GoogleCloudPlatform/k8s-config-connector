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
// proto.message: google.cloud.aiplatform.v1beta1.DeploymentResourcePool
// api.group: vertexai.cnrm.cloud.google.com

package deploymentresourcepool

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vertexAIDeploymentResourcePoolFuzzer())
}

func vertexAIDeploymentResourcePoolFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.DeploymentResourcePool{},
		VertexAIDeploymentResourcePoolSpec_FromProto, VertexAIDeploymentResourcePoolSpec_ToProto,
		VertexAIDeploymentResourcePoolObservedState_FromProto, VertexAIDeploymentResourcePoolObservedState_ToProto,
	)

	f.SpecField(".dedicated_resources")
	f.SpecField(".encryption_spec")
	f.SpecField(".service_account")
	f.SpecField(".disable_container_logging")

	f.Unimplemented_Identity(".name")

	f.StatusField(".create_time")
	f.StatusField(".satisfies_pzs")
	f.StatusField(".satisfies_pzi")

	return f
}
