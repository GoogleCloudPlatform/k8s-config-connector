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

// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.compute.v1

package compute

import (
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeHealthSourceSpec_Sources_FromProto(mapCtx *direct.MapContext, in []string) []computev1beta1.ComputeBackendServiceRef {
	if in == nil {
		return nil
	}
	out := make([]computev1beta1.ComputeBackendServiceRef, 0, len(in))
	for _, s := range in {
		out = append(out, computev1beta1.ComputeBackendServiceRef{External: s})
	}
	return out
}

func ComputeHealthSourceSpec_Sources_ToProto(mapCtx *direct.MapContext, in []computev1beta1.ComputeBackendServiceRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, 0, len(in))
	for _, r := range in {
		if r.External == "" {
			mapCtx.Errorf("reference was not pre-resolved")
		}
		out = append(out, r.External)
	}
	return out
}
