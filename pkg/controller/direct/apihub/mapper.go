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

package apihub

import (
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apihub/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func APIHubInstance_Config_FromProto(mapCtx *direct.MapContext, in *pb.ApiHubInstance_Config) *krm.APIHubInstance_Config {
	if in == nil {
		return nil
	}
	out := &krm.APIHubInstance_Config{}

	if in.GetCmekKeyName() != "" {
		out.CmekKeyRef = &refs.KMSCryptoKeyRef{External: in.GetCmekKeyName()}
	}

	return out
}

func APIHubInstance_Config_ToProto(mapCtx *direct.MapContext, in *krm.APIHubInstance_Config) *pb.ApiHubInstance_Config {
	if in == nil {
		return nil
	}
	out := &pb.ApiHubInstance_Config{}

	if in.CmekKeyRef != nil {
		out.CmekKeyName = in.CmekKeyRef.External
	}

	return out
}
