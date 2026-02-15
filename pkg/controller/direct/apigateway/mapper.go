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

package apigateway

import (
	pb "cloud.google.com/go/apigateway/apiv1/apigatewaypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigateway/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// APIGatewayAPIConfigSpec_ToProto maps the KRM APIConfig to the Proto APIConfig.
// This function is intended to override or augment the generated mapper.
// However, since we cannot redeclare functions in the same package,
// we rely on the generator to NOT generate this function if we define it,
// OR we use a different name and call it from the controller.
// For now, this file is a placeholder for manual mapping logic.

func APIGatewayAPIConfigSpec_ToProto_Manual(mapCtx *direct.MapContext, in *krm.APIGatewayAPIConfigSpec) *pb.ApiConfig {
	if in == nil {
		return nil
	}
	out := &pb.ApiConfig{}
	// Manual mapping logic here if needed.
	// For example, constructing the Name from APIRef + ResourceID?
	// The Name is usually passed in the Create request, not the body.

	return out
}
