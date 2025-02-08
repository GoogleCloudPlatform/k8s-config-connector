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

package servicehealth

import (
	pb "cloud.google.com/go/servicehealth/apiv1/servicehealthpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicehealth/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func Asset_FromProto(mapCtx *direct.MapContext, in *pb.Asset) *krm.Asset {
	if in == nil {
		return nil
	}
	out := &krm.Asset{}
	// MISSING: AssetName
	// MISSING: AssetType
	return out
}
func Asset_ToProto(mapCtx *direct.MapContext, in *krm.Asset) *pb.Asset {
	if in == nil {
		return nil
	}
	out := &pb.Asset{}
	// MISSING: AssetName
	// MISSING: AssetType
	return out
}
func AssetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Asset) *krm.AssetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AssetObservedState{}
	out.AssetName = direct.LazyPtr(in.GetAssetName())
	out.AssetType = direct.LazyPtr(in.GetAssetType())
	return out
}
func AssetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AssetObservedState) *pb.Asset {
	if in == nil {
		return nil
	}
	out := &pb.Asset{}
	out.AssetName = direct.ValueOf(in.AssetName)
	out.AssetType = direct.ValueOf(in.AssetType)
	return out
}
func OrganizationImpact_FromProto(mapCtx *direct.MapContext, in *pb.OrganizationImpact) *krm.OrganizationImpact {
	if in == nil {
		return nil
	}
	out := &krm.OrganizationImpact{}
	// MISSING: Name
	// MISSING: Events
	// MISSING: Asset
	// MISSING: UpdateTime
	return out
}
func OrganizationImpact_ToProto(mapCtx *direct.MapContext, in *krm.OrganizationImpact) *pb.OrganizationImpact {
	if in == nil {
		return nil
	}
	out := &pb.OrganizationImpact{}
	// MISSING: Name
	// MISSING: Events
	// MISSING: Asset
	// MISSING: UpdateTime
	return out
}
func OrganizationImpactObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OrganizationImpact) *krm.OrganizationImpactObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OrganizationImpactObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Events = in.Events
	out.Asset = Asset_FromProto(mapCtx, in.GetAsset())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func OrganizationImpactObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OrganizationImpactObservedState) *pb.OrganizationImpact {
	if in == nil {
		return nil
	}
	out := &pb.OrganizationImpact{}
	out.Name = direct.ValueOf(in.Name)
	out.Events = in.Events
	out.Asset = Asset_ToProto(mapCtx, in.Asset)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func ServicehealthOrganizationImpactObservedState_FromProto(mapCtx *direct.MapContext, in *pb.OrganizationImpact) *krm.ServicehealthOrganizationImpactObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ServicehealthOrganizationImpactObservedState{}
	// MISSING: Name
	// MISSING: Events
	// MISSING: Asset
	// MISSING: UpdateTime
	return out
}
func ServicehealthOrganizationImpactObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ServicehealthOrganizationImpactObservedState) *pb.OrganizationImpact {
	if in == nil {
		return nil
	}
	out := &pb.OrganizationImpact{}
	// MISSING: Name
	// MISSING: Events
	// MISSING: Asset
	// MISSING: UpdateTime
	return out
}
func ServicehealthOrganizationImpactSpec_FromProto(mapCtx *direct.MapContext, in *pb.OrganizationImpact) *krm.ServicehealthOrganizationImpactSpec {
	if in == nil {
		return nil
	}
	out := &krm.ServicehealthOrganizationImpactSpec{}
	// MISSING: Name
	// MISSING: Events
	// MISSING: Asset
	// MISSING: UpdateTime
	return out
}
func ServicehealthOrganizationImpactSpec_ToProto(mapCtx *direct.MapContext, in *krm.ServicehealthOrganizationImpactSpec) *pb.OrganizationImpact {
	if in == nil {
		return nil
	}
	out := &pb.OrganizationImpact{}
	// MISSING: Name
	// MISSING: Events
	// MISSING: Asset
	// MISSING: UpdateTime
	return out
}
