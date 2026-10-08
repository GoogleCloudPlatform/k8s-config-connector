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
// proto.message: google.cloud.networksecurity.v1.InterceptDeployment
// api.group: networksecurity.cnrm.cloud.google.com

package networksecurity

import (
	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(interceptDeploymentFuzzer())
}

func interceptDeploymentFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.InterceptDeployment{},
		NetworkSecurityInterceptDeploymentSpec_v1alpha1_FromProto, NetworkSecurityInterceptDeploymentSpec_v1alpha1_ToProto,
		NetworkSecurityInterceptDeploymentObservedState_v1alpha1_FromProto, NetworkSecurityInterceptDeploymentObservedState_v1alpha1_ToProto,
	)

	f.SpecField(".description")
	f.SpecField(".forwarding_rule")
	f.SpecField(".intercept_deployment_group")
	f.SpecField(".labels")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".state")
	f.StatusField(".reconciling")

	f.Unimplemented_Identity(".name")

	return f
}
