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
// proto.message: google.cloud.apigeeregistry.v1.Api
// api.group: apigeeregistry.cnrm.cloud.google.com

package apigeeregistry

import (
	pb "cloud.google.com/go/apigeeregistry/apiv1/apigeeregistrypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(apigeeRegistryApiFuzzer())
}

func apigeeRegistryApiFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Api{},
		ApigeeRegistryAPISpec_FromProto, ApigeeRegistryAPISpec_ToProto,
		ApigeeRegistryAPIObservedState_FromProto, ApigeeRegistryAPIObservedState_ToProto,
	)

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".availability")
	f.SpecFields.Insert(".recommended_version")
	f.SpecFields.Insert(".recommended_deployment")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".annotations")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")

	f.UnimplementedFields.Insert(".name")

	return f
}
