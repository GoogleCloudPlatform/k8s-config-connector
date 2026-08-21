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

func ComputeHTTPSHealthCheckSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HealthCheck) *krm.ComputeHTTPSHealthCheckSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeHTTPSHealthCheckSpec{}
	out.CheckIntervalSec = in.CheckIntervalSec
	out.Description = in.Description
	out.HealthyThreshold = in.HealthyThreshold
	out.TimeoutSec = in.TimeoutSec
	out.UnhealthyThreshold = in.UnhealthyThreshold

	if in.HttpsHealthCheck != nil {
		out.Host = in.HttpsHealthCheck.Host
		out.Port = in.HttpsHealthCheck.Port
		out.RequestPath = in.HttpsHealthCheck.RequestPath
	}

	return out
}

func ComputeHTTPSHealthCheckSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeHTTPSHealthCheckSpec) *pb.HealthCheck {
	if in == nil {
		return nil
	}
	out := &pb.HealthCheck{}
	out.CheckIntervalSec = in.CheckIntervalSec
	out.Description = in.Description
	out.HealthyThreshold = in.HealthyThreshold
	out.TimeoutSec = in.TimeoutSec
	out.UnhealthyThreshold = in.UnhealthyThreshold

	if in.Host != nil || in.Port != nil || in.RequestPath != nil {
		out.HttpsHealthCheck = &pb.HTTPSHealthCheck{
			Host:        in.Host,
			Port:        in.Port,
			RequestPath: in.RequestPath,
		}
	}

	out.Type = direct.PtrTo("HTTPS")

	return out
}

func ComputeHTTPSHealthCheckStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HealthCheck) *krm.ComputeHTTPSHealthCheckStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeHTTPSHealthCheckStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	out.SelfLink = in.SelfLink
	return out
}

func ComputeHTTPSHealthCheckStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeHTTPSHealthCheckStatus) *pb.HealthCheck {
	if in == nil {
		return nil
	}
	out := &pb.HealthCheck{}
	out.CreationTimestamp = in.CreationTimestamp
	out.SelfLink = in.SelfLink
	return out
}
