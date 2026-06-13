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
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeHTTPHealthCheckStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HTTPHealthCheck) *krm.ComputeHTTPHealthCheckStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeHTTPHealthCheckStatus{}
	// CreationTimestamp and SelfLink are not available in pb.HTTPHealthCheck,
	// so we do not map them here.
	return out
}

func ComputeHTTPHealthCheckStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeHTTPHealthCheckStatus) *pb.HTTPHealthCheck {
	if in == nil {
		return nil
	}
	out := &pb.HTTPHealthCheck{}
	return out
}
