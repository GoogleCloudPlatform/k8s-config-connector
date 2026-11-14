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
	krmcomputev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func GlobalVMExtensionPolicyExtensionPolicy_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.GlobalVmExtensionPolicyExtensionPolicy) *krmcomputev1alpha1.GlobalVMExtensionPolicyExtensionPolicy {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.GlobalVMExtensionPolicyExtensionPolicy{}
	out.PinnedVersion = in.PinnedVersion
	out.StringConfig = in.StringConfig
	return out
}

func GlobalVMExtensionPolicyExtensionPolicy_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.GlobalVMExtensionPolicyExtensionPolicy) *pb.GlobalVmExtensionPolicyExtensionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.GlobalVmExtensionPolicyExtensionPolicy{}
	out.PinnedVersion = in.PinnedVersion
	out.StringConfig = in.StringConfig
	return out
}

func GlobalVMExtensionPolicyLabelSelector_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.GlobalVmExtensionPolicyLabelSelector) *krmcomputev1alpha1.GlobalVMExtensionPolicyLabelSelector {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.GlobalVMExtensionPolicyLabelSelector{}
	out.InclusionLabels = in.InclusionLabels
	return out
}

func GlobalVMExtensionPolicyLabelSelector_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.GlobalVMExtensionPolicyLabelSelector) *pb.GlobalVmExtensionPolicyLabelSelector {
	if in == nil {
		return nil
	}
	out := &pb.GlobalVmExtensionPolicyLabelSelector{}
	out.InclusionLabels = in.InclusionLabels
	return out
}

func GlobalVMExtensionPolicyInstanceSelector_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.GlobalVmExtensionPolicyInstanceSelector) *krmcomputev1alpha1.GlobalVMExtensionPolicyInstanceSelector {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.GlobalVMExtensionPolicyInstanceSelector{}
	out.LabelSelector = GlobalVMExtensionPolicyLabelSelector_v1alpha1_FromProto(mapCtx, in.GetLabelSelector())
	return out
}

func GlobalVMExtensionPolicyInstanceSelector_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.GlobalVMExtensionPolicyInstanceSelector) *pb.GlobalVmExtensionPolicyInstanceSelector {
	if in == nil {
		return nil
	}
	out := &pb.GlobalVmExtensionPolicyInstanceSelector{}
	out.LabelSelector = GlobalVMExtensionPolicyLabelSelector_v1alpha1_ToProto(mapCtx, in.LabelSelector)
	return out
}

func GlobalExtensionPolicies_FromProto(mapCtx *direct.MapContext, in map[string]*pb.GlobalVmExtensionPolicyExtensionPolicy) map[string]krmcomputev1alpha1.GlobalVMExtensionPolicyExtensionPolicy {
	if in == nil {
		return nil
	}
	out := make(map[string]krmcomputev1alpha1.GlobalVMExtensionPolicyExtensionPolicy, len(in))
	for k, v := range in {
		if mapped := GlobalVMExtensionPolicyExtensionPolicy_v1alpha1_FromProto(mapCtx, v); mapped != nil {
			out[k] = *mapped
		}
	}
	return out
}

func GlobalExtensionPolicies_ToProto(mapCtx *direct.MapContext, in map[string]krmcomputev1alpha1.GlobalVMExtensionPolicyExtensionPolicy) map[string]*pb.GlobalVmExtensionPolicyExtensionPolicy {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.GlobalVmExtensionPolicyExtensionPolicy, len(in))
	for k, v := range in {
		if mapped := GlobalVMExtensionPolicyExtensionPolicy_v1alpha1_ToProto(mapCtx, &v); mapped != nil {
			out[k] = mapped
		}
	}
	return out
}

func GlobalVMExtensionPolicyRolloutOperationRolloutInput_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.GlobalVmExtensionPolicyRolloutOperationRolloutInput) *krmcomputev1alpha1.GlobalVMExtensionPolicyRolloutOperationRolloutInput {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.GlobalVMExtensionPolicyRolloutOperationRolloutInput{}
	out.ConflictBehavior = in.ConflictBehavior
	out.Name = in.Name
	out.PredefinedRolloutPlan = in.PredefinedRolloutPlan
	out.RetryUuid = in.RetryUuid
	return out
}

func GlobalVMExtensionPolicyRolloutOperationRolloutInput_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.GlobalVMExtensionPolicyRolloutOperationRolloutInput) *pb.GlobalVmExtensionPolicyRolloutOperationRolloutInput {
	if in == nil {
		return nil
	}
	out := &pb.GlobalVmExtensionPolicyRolloutOperationRolloutInput{}
	out.ConflictBehavior = in.ConflictBehavior
	out.Name = in.Name
	out.PredefinedRolloutPlan = in.PredefinedRolloutPlan
	out.RetryUuid = in.RetryUuid
	return out
}

func ComputeGlobalVMExtensionPolicySpec_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.GlobalVmExtensionPolicy) *krmcomputev1alpha1.ComputeGlobalVMExtensionPolicySpec {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.ComputeGlobalVMExtensionPolicySpec{}
	out.Description = in.Description
	out.ExtensionPolicies = GlobalExtensionPolicies_FromProto(mapCtx, in.ExtensionPolicies)
	out.InstanceSelectors = direct.Slice_FromProto(mapCtx, in.InstanceSelectors, GlobalVMExtensionPolicyInstanceSelector_v1alpha1_FromProto)
	// MISSING: Name
	out.Priority = in.Priority
	if in.RolloutOperation != nil && in.RolloutOperation.RolloutInput != nil {
		out.RolloutOperation = &krmcomputev1alpha1.GlobalVMExtensionPolicyRolloutOperation{
			RolloutInput: GlobalVMExtensionPolicyRolloutOperationRolloutInput_v1alpha1_FromProto(mapCtx, in.RolloutOperation.RolloutInput),
		}
	}
	return out
}

func ComputeGlobalVMExtensionPolicySpec_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.ComputeGlobalVMExtensionPolicySpec) *pb.GlobalVmExtensionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.GlobalVmExtensionPolicy{}
	out.Description = in.Description
	out.ExtensionPolicies = GlobalExtensionPolicies_ToProto(mapCtx, in.ExtensionPolicies)
	out.InstanceSelectors = direct.Slice_ToProto(mapCtx, in.InstanceSelectors, GlobalVMExtensionPolicyInstanceSelector_v1alpha1_ToProto)
	// MISSING: Name
	out.Priority = in.Priority
	if in.RolloutOperation != nil && in.RolloutOperation.RolloutInput != nil {
		out.RolloutOperation = &pb.GlobalVmExtensionPolicyRolloutOperation{
			RolloutInput: GlobalVMExtensionPolicyRolloutOperationRolloutInput_v1alpha1_ToProto(mapCtx, in.RolloutOperation.RolloutInput),
		}
	}
	return out
}

func ComputeGlobalVMExtensionPolicyObservedState_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.GlobalVmExtensionPolicy) *krmcomputev1alpha1.ComputeGlobalVMExtensionPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krmcomputev1alpha1.ComputeGlobalVMExtensionPolicyObservedState{}
	out.CreationTimestamp = in.CreationTimestamp
	out.ID = in.Id
	out.Kind = in.Kind
	out.ScopedResourceStatus = in.ScopedResourceStatus
	out.SelfLink = in.SelfLink
	out.SelfLinkWithID = in.SelfLinkWithId
	out.UpdateTimestamp = in.UpdateTimestamp

	if in.RolloutOperation != nil && in.RolloutOperation.RolloutStatus != nil {
		status := in.RolloutOperation.RolloutStatus
		obsStatus := &krmcomputev1alpha1.GlobalVMExtensionPolicyRolloutOperationRolloutStatus{}
		for _, r := range status.CurrentRollouts {
			meta := krmcomputev1alpha1.GlobalVMExtensionPolicyRolloutOperationRolloutStatusRolloutMetadata{
				Rollout:     r.Rollout,
				RolloutPlan: r.RolloutPlan,
				State:       r.State,
			}
			if r.LocationRolloutStatus != nil {
				meta.LocationRolloutStatus = make(map[string]krmcomputev1alpha1.GlobalVMExtensionPolicyRolloutOperationRolloutStatusRolloutMetadataLocationRolloutStatus, len(r.LocationRolloutStatus))
				for loc, lStatus := range r.LocationRolloutStatus {
					meta.LocationRolloutStatus[loc] = krmcomputev1alpha1.GlobalVMExtensionPolicyRolloutOperationRolloutStatusRolloutMetadataLocationRolloutStatus{
						State: lStatus.State,
					}
				}
			}
			obsStatus.CurrentRollouts = append(obsStatus.CurrentRollouts, meta)
		}
		if status.PreviousRollout != nil {
			r := status.PreviousRollout
			meta := &krmcomputev1alpha1.GlobalVMExtensionPolicyRolloutOperationRolloutStatusRolloutMetadata{
				Rollout:     r.Rollout,
				RolloutPlan: r.RolloutPlan,
				State:       r.State,
			}
			if r.LocationRolloutStatus != nil {
				meta.LocationRolloutStatus = make(map[string]krmcomputev1alpha1.GlobalVMExtensionPolicyRolloutOperationRolloutStatusRolloutMetadataLocationRolloutStatus, len(r.LocationRolloutStatus))
				for loc, lStatus := range r.LocationRolloutStatus {
					meta.LocationRolloutStatus[loc] = krmcomputev1alpha1.GlobalVMExtensionPolicyRolloutOperationRolloutStatusRolloutMetadataLocationRolloutStatus{
						State: lStatus.State,
					}
				}
			}
			obsStatus.PreviousRollout = meta
		}
		out.RolloutOperation = &krmcomputev1alpha1.GlobalVMExtensionPolicyRolloutOperationObservedState{
			RolloutStatus: obsStatus,
		}
	}
	return out
}

func ComputeGlobalVMExtensionPolicyObservedState_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmcomputev1alpha1.ComputeGlobalVMExtensionPolicyObservedState) *pb.GlobalVmExtensionPolicy {
	if in == nil {
		return nil
	}
	out := &pb.GlobalVmExtensionPolicy{}
	out.CreationTimestamp = in.CreationTimestamp
	out.Id = in.ID
	out.Kind = in.Kind
	out.ScopedResourceStatus = in.ScopedResourceStatus
	out.SelfLink = in.SelfLink
	out.SelfLinkWithId = in.SelfLinkWithID
	out.UpdateTimestamp = in.UpdateTimestamp

	if in.RolloutOperation != nil && in.RolloutOperation.RolloutStatus != nil {
		status := in.RolloutOperation.RolloutStatus
		pbStatus := &pb.GlobalVmExtensionPolicyRolloutOperationRolloutStatus{}
		for _, r := range status.CurrentRollouts {
			meta := pb.GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadata{
				Rollout:     r.Rollout,
				RolloutPlan: r.RolloutPlan,
				State:       r.State,
			}
			if r.LocationRolloutStatus != nil {
				meta.LocationRolloutStatus = make(map[string]*pb.GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadataLocationRolloutStatus, len(r.LocationRolloutStatus))
				for loc, lStatus := range r.LocationRolloutStatus {
					meta.LocationRolloutStatus[loc] = &pb.GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadataLocationRolloutStatus{
						State: lStatus.State,
					}
				}
			}
			pbStatus.CurrentRollouts = append(pbStatus.CurrentRollouts, &meta)
		}
		if status.PreviousRollout != nil {
			r := status.PreviousRollout
			meta := &pb.GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadata{
				Rollout:     r.Rollout,
				RolloutPlan: r.RolloutPlan,
				State:       r.State,
			}
			if r.LocationRolloutStatus != nil {
				meta.LocationRolloutStatus = make(map[string]*pb.GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadataLocationRolloutStatus, len(r.LocationRolloutStatus))
				for loc, lStatus := range r.LocationRolloutStatus {
					meta.LocationRolloutStatus[loc] = &pb.GlobalVmExtensionPolicyRolloutOperationRolloutStatusRolloutMetadataLocationRolloutStatus{
						State: lStatus.State,
					}
				}
			}
			pbStatus.PreviousRollout = meta
		}
		out.RolloutOperation = &pb.GlobalVmExtensionPolicyRolloutOperation{
			RolloutStatus: pbStatus,
		}
	}
	return out
}
