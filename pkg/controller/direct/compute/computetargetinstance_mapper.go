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

// ComputeTargetInstanceSpec_v1beta1_FromProto maps a pb.TargetInstance proto representation to a krm.ComputeTargetInstanceSpec.
// We hand-code this mapping because Zone in the GCP proto is a *string, whereas in the KRM spec it is a string.
// Additionally, InstanceRef is a pointer to InstanceRef in the KRM spec but mapped as a string pointer in the GCP proto.
func ComputeTargetInstanceSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.TargetInstance) *krm.ComputeTargetInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeTargetInstanceSpec{}
	out.Description = in.Description
	if in.GetInstance() != "" {
		out.InstanceRef = &krm.InstanceRef{External: in.GetInstance()}
	}
	out.NATPolicy = in.NatPolicy
	if in.GetNetwork() != "" {
		out.NetworkRef = &krm.ComputeNetworkRef{External: in.GetNetwork()}
	}
	if in.GetSecurityPolicy() != "" {
		out.SecurityPolicyRef = &krm.ComputeSecurityPolicyRef{External: in.GetSecurityPolicy()}
	}
	out.Zone = in.GetZone()
	return out
}

// ComputeTargetInstanceSpec_v1beta1_ToProto maps a krm.ComputeTargetInstanceSpec to a pb.TargetInstance proto representation.
// We hand-code this mapping because Zone in the GCP proto is a *string, whereas in the KRM spec it is a string.
// Additionally, InstanceRef is a pointer to InstanceRef in the KRM spec but mapped as a string pointer in the GCP proto.
func ComputeTargetInstanceSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeTargetInstanceSpec) *pb.TargetInstance {
	if in == nil {
		return nil
	}
	out := &pb.TargetInstance{}
	out.Description = in.Description
	if in.InstanceRef != nil {
		out.Instance = &in.InstanceRef.External
	}
	out.NatPolicy = in.NATPolicy
	if in.NetworkRef != nil {
		out.Network = &in.NetworkRef.External
	}
	if in.SecurityPolicyRef != nil {
		out.SecurityPolicy = &in.SecurityPolicyRef.External
	}
	if in.Zone != "" {
		out.Zone = &in.Zone
	}
	return out
}
