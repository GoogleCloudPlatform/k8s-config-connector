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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeURLMapDefaultService_v1beta1_FromProto(mapCtx *direct.MapContext, in *string) *krm.ComputeURLMapDefaultService {
	if in == nil || *in == "" {
		return nil
	}
	out := &krm.ComputeURLMapDefaultService{}
	if strings.Contains(*in, "/backendBuckets/") {
		out.BackendBucketRef = &v1alpha1.ResourceRef{External: *in}
	} else {
		out.BackendServiceRef = &v1alpha1.ResourceRef{External: *in}
	}
	return out
}

func ComputeURLMapDefaultService_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeURLMapDefaultService) *string {
	if in == nil {
		return nil
	}
	if in.BackendBucketRef != nil && in.BackendBucketRef.External != "" {
		return &in.BackendBucketRef.External
	}
	if in.BackendServiceRef != nil && in.BackendServiceRef.External != "" {
		return &in.BackendServiceRef.External
	}
	return nil
}
