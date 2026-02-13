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

package compute

import (
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeURLMapServiceRef_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeURLMapServiceRef) *string {
	if in == nil {
		return nil
	}
	if in.BackendBucketRef != nil {
		return &in.BackendBucketRef.External
	}
	if in.BackendServiceRef != nil {
		return &in.BackendServiceRef.External
	}
	return nil
}

func ComputeURLMapServiceRef_v1beta1_FromProto(mapCtx *direct.MapContext, in *string) *krm.ComputeURLMapServiceRef {
	if in == nil || *in == "" {
		return nil
	}

	res := &krm.ComputeURLMapServiceRef{}
	// Simple string check
	s := *in
	if strings.Contains(s, "/backendBuckets/") {
		res.BackendBucketRef = &krm.ComputeBackendBucketRef{External: s}
	} else {
		res.BackendServiceRef = &krm.ComputeBackendServiceRef{External: s}
	}
	return res
}
