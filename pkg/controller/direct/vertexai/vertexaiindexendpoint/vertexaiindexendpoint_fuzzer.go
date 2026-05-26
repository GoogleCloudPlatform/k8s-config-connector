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

package vertexaiindexendpoint

import (
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	api "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/vertexai"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(vertexaiindexendpointFuzzer())
}

func vertexaiindexendpointFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.IndexEndpoint{},
		api.VertexAIIndexEndpointSpec_FromProto,
		api.VertexAIIndexEndpointSpec_ToProto,
		api.VertexAIIndexEndpointObservedState_FromProto,
		api.VertexAIIndexEndpointObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name") // special field

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".network")
	f.SpecFields.Insert(".enable_private_service_connect")
	f.SpecFields.Insert(".private_service_connect_config")
	f.SpecFields.Insert(".public_endpoint_enabled")
	f.SpecFields.Insert(".encryption_spec")

	f.StatusFields.Insert(".deployed_indexes")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".public_endpoint_domain_name")
	f.StatusFields.Insert(".satisfies_pzs")
	f.StatusFields.Insert(".satisfies_pzi")

	f.Unimplemented_Etag()

	return f
}
