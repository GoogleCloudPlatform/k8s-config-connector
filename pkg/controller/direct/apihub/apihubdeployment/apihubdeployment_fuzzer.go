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

package apihubdeployment

import (
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/apihub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(apihubDeploymentFuzzer())
}

func apihubDeploymentFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Deployment{},
		apihub.APIHubDeploymentSpec_FromProto, apihub.APIHubDeploymentSpec_ToProto,
		apihub.APIHubDeploymentObservedState_FromProto, apihub.APIHubDeploymentObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".attributes")
	f.UnimplementedFields.Insert(".source_metadata")
	f.UnimplementedFields.Insert(".management_url")
	f.UnimplementedFields.Insert(".source_uri")
	f.UnimplementedFields.Insert(".source_project")
	f.UnimplementedFields.Insert(".source_environment")

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".documentation")
	f.SpecFields.Insert(".deployment_type")
	f.SpecFields.Insert(".resource_uri")
	f.SpecFields.Insert(".endpoints")
	f.SpecFields.Insert(".slo")
	f.SpecFields.Insert(".environment")

	f.UnimplementedFields.Insert(".deployment_type")
	f.UnimplementedFields.Insert(".slo")
	f.UnimplementedFields.Insert(".environment")

	f.StatusFields.Insert(".api_versions")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")

	return f
}
