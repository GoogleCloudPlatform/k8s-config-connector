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

package configdeployment

import (
	pb "cloud.google.com/go/config/apiv1/configpb"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(fuzzConfigDeploymentGroup())
}

func fuzzConfigDeploymentGroup() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.DeploymentGroup{},
		ConfigDeploymentGroupSpec_FromProto, ConfigDeploymentGroupSpec_ToProto,
		ConfigDeploymentGroupObservedState_FromProto, ConfigDeploymentGroupObservedState_ToProto,
	)

	f.Unimplemented_Identity(".name")

	f.SpecField(".labels")
	f.SpecField(".annotations")
	f.SpecField(".deployment_units")
	f.SpecField(".deployment_units.id")
	f.SpecField(".deployment_units.deployment")
	f.SpecField(".deployment_units.dependencies")

	f.StatusField(".create_time")
	f.StatusField(".update_time")
	f.StatusField(".state")
	f.StatusField(".state_description")
	f.StatusField(".provisioning_state")
	f.StatusField(".provisioning_state_description")
	f.StatusField(".provisioning_error")
	f.StatusField(".provisioning_error.code")
	f.StatusField(".provisioning_error.message")
	f.Unimplemented_NotYetTriaged(".provisioning_error.details")

	return f
}
