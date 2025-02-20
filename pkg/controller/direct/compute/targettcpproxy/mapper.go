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

package targettcpproxy

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeTargetTCPProxySpec_BackendServiceRef_FromProto(mapCtx *direct.MapContext, in string) *krm.ComputeBackendServiceRef {
	if in == "" {
		return nil
	}
	return &krm.ComputeBackendServiceRef{
		External: in,
	}
}

func ComputeTargetTCPProxySpec_BackendServiceRef_ToProto(mapCtx *direct.MapContext, in *krm.ComputeBackendServiceRef) *string {
	if in == nil {
		return nil
	}
	if in.External == "" {
		mapCtx.Errorf("reference %s was not pre-resolved", in.Name)
	}
	return direct.LazyPtr(in.External)
}
