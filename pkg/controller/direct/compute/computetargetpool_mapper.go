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

func ComputeTargetPoolSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.TargetPool) *krm.ComputeTargetPoolSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeTargetPoolSpec{}

	if in.BackupPool != nil && *in.BackupPool != "" {
		out.BackupTargetPoolRef = &krm.TargetPoolResourceRef{External: *in.BackupPool}
	}

	out.Description = in.Description

	if in.FailoverRatio != nil {
		val := float64(*in.FailoverRatio)
		out.FailoverRatio = &val
	}

	if len(in.HealthChecks) > 0 {
		out.HealthChecks = make([]krm.TargetpoolHealthChecks, len(in.HealthChecks))
		for i, hc := range in.HealthChecks {
			out.HealthChecks[i] = krm.TargetpoolHealthChecks{
				HttpHealthCheckRef: &krm.TargetPoolResourceRef{External: hc},
			}
		}
	}

	if len(in.Instances) > 0 {
		out.Instances = make([]krm.TargetPoolResourceRef, len(in.Instances))
		for i, inst := range in.Instances {
			out.Instances[i] = krm.TargetPoolResourceRef{External: inst}
		}
	}

	if in.Region != nil {
		out.Region = *in.Region
	}

	if in.SecurityPolicy != nil && *in.SecurityPolicy != "" {
		out.SecurityPolicyRef = &krm.TargetPoolResourceRef{External: *in.SecurityPolicy}
	}

	out.SessionAffinity = in.SessionAffinity

	return out
}

func ComputeTargetPoolSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeTargetPoolSpec) *pb.TargetPool {
	if in == nil {
		return nil
	}
	out := &pb.TargetPool{}

	if in.BackupTargetPoolRef != nil {
		out.BackupPool = direct.LazyPtr(in.BackupTargetPoolRef.External)
	}

	out.Description = in.Description

	if in.FailoverRatio != nil {
		val := float32(*in.FailoverRatio)
		out.FailoverRatio = &val
	}

	if len(in.HealthChecks) > 0 {
		out.HealthChecks = make([]string, len(in.HealthChecks))
		for i, hc := range in.HealthChecks {
			if hc.HttpHealthCheckRef != nil {
				out.HealthChecks[i] = hc.HttpHealthCheckRef.External
			}
		}
	}

	if len(in.Instances) > 0 {
		out.Instances = make([]string, len(in.Instances))
		for i, inst := range in.Instances {
			out.Instances[i] = inst.External
		}
	}

	if in.Region != "" {
		out.Region = direct.LazyPtr(in.Region)
	}

	if in.SecurityPolicyRef != nil {
		out.SecurityPolicy = direct.LazyPtr(in.SecurityPolicyRef.External)
	}

	out.SessionAffinity = in.SessionAffinity

	return out
}
