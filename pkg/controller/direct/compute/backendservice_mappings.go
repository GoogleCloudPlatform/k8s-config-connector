// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compute

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeBackendServiceSpec_HealthChecks_FromProto(mapCtx *direct.MapContext, in []string) []krm.BackendserviceHealthChecks {
	if len(in) == 0 {
		return nil
	}
	var out []krm.BackendserviceHealthChecks
	for _, hc := range in {
		out = append(out, krm.BackendserviceHealthChecks{
			HealthCheckRef: &refsv1beta1.ComputeHealthCheckRef{
				External: hc,
			},
		})
	}
	return out
}

func ComputeBackendServiceSpec_HealthChecks_ToProto(mapCtx *direct.MapContext, in []krm.BackendserviceHealthChecks) []string {
	if len(in) == 0 {
		return nil
	}
	var out []string
	for _, hc := range in {
		if hc.HealthCheckRef != nil && hc.HealthCheckRef.External != "" {
			out = append(out, hc.HealthCheckRef.External)
		} else if hc.HttpHealthCheckRef != nil && hc.HttpHealthCheckRef.External != "" {
			out = append(out, hc.HttpHealthCheckRef.External)
		}
	}
	return out
}
