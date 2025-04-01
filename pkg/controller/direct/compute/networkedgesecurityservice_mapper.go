// Copyright 2025 Google LLC
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

// +generated:mapper
// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.compute.v1

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeNetworkEdgeSecurityServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.NetworkEdgeSecurityService) *krm.ComputeNetworkEdgeSecurityServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ComputeNetworkEdgeSecurityServiceObservedState{}
	out.CreationTimestamp = in.CreationTimestamp
	out.ID = in.Id
	out.Kind = in.Kind
	// MISSING: Name
	out.Region = in.Region
	out.SelfLink = in.SelfLink
	out.SelfLinkWithID = in.SelfLinkWithId
	return out
}
func ComputeNetworkEdgeSecurityServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ComputeNetworkEdgeSecurityServiceObservedState) *pb.NetworkEdgeSecurityService {
	if in == nil {
		return nil
	}
	out := &pb.NetworkEdgeSecurityService{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Id = in.ID
	out.Kind = in.Kind
	// MISSING: Name
	out.Region = in.Region
	out.SelfLink = in.SelfLink
	out.SelfLinkWithId = in.SelfLinkWithID
	return out
}
func ComputeNetworkEdgeSecurityServiceSpec_FromProto(mapCtx *direct.MapContext, in *pb.NetworkEdgeSecurityService) *krm.ComputeNetworkEdgeSecurityServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeNetworkEdgeSecurityServiceSpec{}
	out.Description = in.Description
	out.Fingerprint = in.Fingerprint
	// MISSING: Name
	out.SecurityPolicy = in.SecurityPolicy
	return out
}
func ComputeNetworkEdgeSecurityServiceSpec_ToProto(mapCtx *direct.MapContext, in *krm.ComputeNetworkEdgeSecurityServiceSpec) *pb.NetworkEdgeSecurityService {
	if in == nil {
		return nil
	}
	out := &pb.NetworkEdgeSecurityService{}
	out.Description = in.Description
	out.Fingerprint = in.Fingerprint
	// MISSING: Name
	out.SecurityPolicy = in.SecurityPolicy
	return out
}
