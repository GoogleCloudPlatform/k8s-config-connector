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

package commerce

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/commerce/consumer/procurement/apiv1/procurementpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/commerce/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AssignmentProtocol_FromProto(mapCtx *direct.MapContext, in *pb.AssignmentProtocol) *krm.AssignmentProtocol {
	if in == nil {
		return nil
	}
	out := &krm.AssignmentProtocol{}
	out.ManualAssignmentType = AssignmentProtocol_ManualAssignmentType_FromProto(mapCtx, in.GetManualAssignmentType())
	out.AutoAssignmentType = AssignmentProtocol_AutoAssignmentType_FromProto(mapCtx, in.GetAutoAssignmentType())
	return out
}
func AssignmentProtocol_ToProto(mapCtx *direct.MapContext, in *krm.AssignmentProtocol) *pb.AssignmentProtocol {
	if in == nil {
		return nil
	}
	out := &pb.AssignmentProtocol{}
	if oneof := AssignmentProtocol_ManualAssignmentType_ToProto(mapCtx, in.ManualAssignmentType); oneof != nil {
		out.AssignmentType = &pb.AssignmentProtocol_ManualAssignmentType_{ManualAssignmentType: oneof}
	}
	if oneof := AssignmentProtocol_AutoAssignmentType_ToProto(mapCtx, in.AutoAssignmentType); oneof != nil {
		out.AssignmentType = &pb.AssignmentProtocol_AutoAssignmentType_{AutoAssignmentType: oneof}
	}
	return out
}
func AssignmentProtocol_AutoAssignmentType_FromProto(mapCtx *direct.MapContext, in *pb.AssignmentProtocol_AutoAssignmentType) *krm.AssignmentProtocol_AutoAssignmentType {
	if in == nil {
		return nil
	}
	out := &krm.AssignmentProtocol_AutoAssignmentType{}
	out.InactiveLicenseTtl = direct.StringDuration_FromProto(mapCtx, in.GetInactiveLicenseTtl())
	return out
}
func AssignmentProtocol_AutoAssignmentType_ToProto(mapCtx *direct.MapContext, in *krm.AssignmentProtocol_AutoAssignmentType) *pb.AssignmentProtocol_AutoAssignmentType {
	if in == nil {
		return nil
	}
	out := &pb.AssignmentProtocol_AutoAssignmentType{}
	out.InactiveLicenseTtl = direct.StringDuration_ToProto(mapCtx, in.InactiveLicenseTtl)
	return out
}
func AssignmentProtocol_ManualAssignmentType_FromProto(mapCtx *direct.MapContext, in *pb.AssignmentProtocol_ManualAssignmentType) *krm.AssignmentProtocol_ManualAssignmentType {
	if in == nil {
		return nil
	}
	out := &krm.AssignmentProtocol_ManualAssignmentType{}
	return out
}
func AssignmentProtocol_ManualAssignmentType_ToProto(mapCtx *direct.MapContext, in *krm.AssignmentProtocol_ManualAssignmentType) *pb.AssignmentProtocol_ManualAssignmentType {
	if in == nil {
		return nil
	}
	out := &pb.AssignmentProtocol_ManualAssignmentType{}
	return out
}
func CommerceLicensePoolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LicensePool) *krm.CommerceLicensePoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CommerceLicensePoolObservedState{}
	// MISSING: Name
	// MISSING: LicenseAssignmentProtocol
	// MISSING: AvailableLicenseCount
	// MISSING: TotalLicenseCount
	return out
}
func CommerceLicensePoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CommerceLicensePoolObservedState) *pb.LicensePool {
	if in == nil {
		return nil
	}
	out := &pb.LicensePool{}
	// MISSING: Name
	// MISSING: LicenseAssignmentProtocol
	// MISSING: AvailableLicenseCount
	// MISSING: TotalLicenseCount
	return out
}
func CommerceLicensePoolSpec_FromProto(mapCtx *direct.MapContext, in *pb.LicensePool) *krm.CommerceLicensePoolSpec {
	if in == nil {
		return nil
	}
	out := &krm.CommerceLicensePoolSpec{}
	// MISSING: Name
	// MISSING: LicenseAssignmentProtocol
	// MISSING: AvailableLicenseCount
	// MISSING: TotalLicenseCount
	return out
}
func CommerceLicensePoolSpec_ToProto(mapCtx *direct.MapContext, in *krm.CommerceLicensePoolSpec) *pb.LicensePool {
	if in == nil {
		return nil
	}
	out := &pb.LicensePool{}
	// MISSING: Name
	// MISSING: LicenseAssignmentProtocol
	// MISSING: AvailableLicenseCount
	// MISSING: TotalLicenseCount
	return out
}
func LicensePool_FromProto(mapCtx *direct.MapContext, in *pb.LicensePool) *krm.LicensePool {
	if in == nil {
		return nil
	}
	out := &krm.LicensePool{}
	out.Name = direct.LazyPtr(in.GetName())
	out.LicenseAssignmentProtocol = AssignmentProtocol_FromProto(mapCtx, in.GetLicenseAssignmentProtocol())
	// MISSING: AvailableLicenseCount
	// MISSING: TotalLicenseCount
	return out
}
func LicensePool_ToProto(mapCtx *direct.MapContext, in *krm.LicensePool) *pb.LicensePool {
	if in == nil {
		return nil
	}
	out := &pb.LicensePool{}
	out.Name = direct.ValueOf(in.Name)
	out.LicenseAssignmentProtocol = AssignmentProtocol_ToProto(mapCtx, in.LicenseAssignmentProtocol)
	// MISSING: AvailableLicenseCount
	// MISSING: TotalLicenseCount
	return out
}
func LicensePoolObservedState_FromProto(mapCtx *direct.MapContext, in *pb.LicensePool) *krm.LicensePoolObservedState {
	if in == nil {
		return nil
	}
	out := &krm.LicensePoolObservedState{}
	// MISSING: Name
	// MISSING: LicenseAssignmentProtocol
	out.AvailableLicenseCount = direct.LazyPtr(in.GetAvailableLicenseCount())
	out.TotalLicenseCount = direct.LazyPtr(in.GetTotalLicenseCount())
	return out
}
func LicensePoolObservedState_ToProto(mapCtx *direct.MapContext, in *krm.LicensePoolObservedState) *pb.LicensePool {
	if in == nil {
		return nil
	}
	out := &pb.LicensePool{}
	// MISSING: Name
	// MISSING: LicenseAssignmentProtocol
	out.AvailableLicenseCount = direct.ValueOf(in.AvailableLicenseCount)
	out.TotalLicenseCount = direct.ValueOf(in.TotalLicenseCount)
	return out
}
