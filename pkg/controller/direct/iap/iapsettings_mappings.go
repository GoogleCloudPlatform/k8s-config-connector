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

package iap

import (
	pb "cloud.google.com/go/iap/apiv1/iappb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iap/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AccessDeniedPageSettings_ToProto(mapCtx *direct.MapContext, in *krm.AccessDeniedPageSettings) *pb.AccessDeniedPageSettings {
	if in == nil {
		return nil
	}
	out := &pb.AccessDeniedPageSettings{}
	out.AccessDeniedPageUri = direct.StringValue_ToProto(mapCtx, in.AccessDeniedPageURI)
	out.GenerateTroubleshootingUri = direct.BoolValue_ToProto(mapCtx, in.GenerateTroubleshootingURI)
	out.RemediationTokenGenerationEnabled = direct.BoolValue_ToProto(mapCtx, in.RemediationTokenGenerationEnabled) // this line is manually edited because proto field is incorrectly marked as oneof
	return out
}
