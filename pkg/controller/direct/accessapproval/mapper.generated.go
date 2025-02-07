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

package accessapproval

import (
	pb "cloud.google.com/go/accessapproval/apiv1/accessapprovalpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/accessapproval/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func AccessApprovalSettings_FromProto(mapCtx *direct.MapContext, in *pb.AccessApprovalSettings) *krm.AccessApprovalSettings {
	if in == nil {
		return nil
	}
	out := &krm.AccessApprovalSettings{}
	out.Name = direct.LazyPtr(in.GetName())
	out.NotificationEmails = in.NotificationEmails
	out.EnrolledServices = direct.Slice_FromProto(mapCtx, in.EnrolledServices, EnrolledService_FromProto)
	// MISSING: EnrolledAncestor
	out.ActiveKeyVersion = direct.LazyPtr(in.GetActiveKeyVersion())
	// MISSING: AncestorHasActiveKeyVersion
	// MISSING: InvalidKeyVersion
	return out
}
func AccessApprovalSettings_ToProto(mapCtx *direct.MapContext, in *krm.AccessApprovalSettings) *pb.AccessApprovalSettings {
	if in == nil {
		return nil
	}
	out := &pb.AccessApprovalSettings{}
	out.Name = direct.ValueOf(in.Name)
	out.NotificationEmails = in.NotificationEmails
	out.EnrolledServices = direct.Slice_ToProto(mapCtx, in.EnrolledServices, EnrolledService_ToProto)
	// MISSING: EnrolledAncestor
	out.ActiveKeyVersion = direct.ValueOf(in.ActiveKeyVersion)
	// MISSING: AncestorHasActiveKeyVersion
	// MISSING: InvalidKeyVersion
	return out
}
func AccessApprovalSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccessApprovalSettings) *krm.AccessApprovalSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AccessApprovalSettingsObservedState{}
	// MISSING: Name
	// MISSING: NotificationEmails
	// MISSING: EnrolledServices
	out.EnrolledAncestor = direct.LazyPtr(in.GetEnrolledAncestor())
	// MISSING: ActiveKeyVersion
	out.AncestorHasActiveKeyVersion = direct.LazyPtr(in.GetAncestorHasActiveKeyVersion())
	out.InvalidKeyVersion = direct.LazyPtr(in.GetInvalidKeyVersion())
	return out
}
func AccessApprovalSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AccessApprovalSettingsObservedState) *pb.AccessApprovalSettings {
	if in == nil {
		return nil
	}
	out := &pb.AccessApprovalSettings{}
	// MISSING: Name
	// MISSING: NotificationEmails
	// MISSING: EnrolledServices
	out.EnrolledAncestor = direct.ValueOf(in.EnrolledAncestor)
	// MISSING: ActiveKeyVersion
	out.AncestorHasActiveKeyVersion = direct.ValueOf(in.AncestorHasActiveKeyVersion)
	out.InvalidKeyVersion = direct.ValueOf(in.InvalidKeyVersion)
	return out
}
func AccessapprovalAccessApprovalSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccessApprovalSettings) *krm.AccessapprovalAccessApprovalSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AccessapprovalAccessApprovalSettingsObservedState{}
	// MISSING: Name
	// MISSING: NotificationEmails
	// MISSING: EnrolledServices
	// MISSING: EnrolledAncestor
	// MISSING: ActiveKeyVersion
	// MISSING: AncestorHasActiveKeyVersion
	// MISSING: InvalidKeyVersion
	return out
}
func AccessapprovalAccessApprovalSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AccessapprovalAccessApprovalSettingsObservedState) *pb.AccessApprovalSettings {
	if in == nil {
		return nil
	}
	out := &pb.AccessApprovalSettings{}
	// MISSING: Name
	// MISSING: NotificationEmails
	// MISSING: EnrolledServices
	// MISSING: EnrolledAncestor
	// MISSING: ActiveKeyVersion
	// MISSING: AncestorHasActiveKeyVersion
	// MISSING: InvalidKeyVersion
	return out
}
func AccessapprovalAccessApprovalSettingsSpec_FromProto(mapCtx *direct.MapContext, in *pb.AccessApprovalSettings) *krm.AccessapprovalAccessApprovalSettingsSpec {
	if in == nil {
		return nil
	}
	out := &krm.AccessapprovalAccessApprovalSettingsSpec{}
	// MISSING: Name
	// MISSING: NotificationEmails
	// MISSING: EnrolledServices
	// MISSING: EnrolledAncestor
	// MISSING: ActiveKeyVersion
	// MISSING: AncestorHasActiveKeyVersion
	// MISSING: InvalidKeyVersion
	return out
}
func AccessapprovalAccessApprovalSettingsSpec_ToProto(mapCtx *direct.MapContext, in *krm.AccessapprovalAccessApprovalSettingsSpec) *pb.AccessApprovalSettings {
	if in == nil {
		return nil
	}
	out := &pb.AccessApprovalSettings{}
	// MISSING: Name
	// MISSING: NotificationEmails
	// MISSING: EnrolledServices
	// MISSING: EnrolledAncestor
	// MISSING: ActiveKeyVersion
	// MISSING: AncestorHasActiveKeyVersion
	// MISSING: InvalidKeyVersion
	return out
}
func EnrolledService_FromProto(mapCtx *direct.MapContext, in *pb.EnrolledService) *krm.EnrolledService {
	if in == nil {
		return nil
	}
	out := &krm.EnrolledService{}
	out.CloudProduct = direct.LazyPtr(in.GetCloudProduct())
	out.EnrollmentLevel = direct.Enum_FromProto(mapCtx, in.GetEnrollmentLevel())
	return out
}
func EnrolledService_ToProto(mapCtx *direct.MapContext, in *krm.EnrolledService) *pb.EnrolledService {
	if in == nil {
		return nil
	}
	out := &pb.EnrolledService{}
	out.CloudProduct = direct.ValueOf(in.CloudProduct)
	out.EnrollmentLevel = direct.Enum_ToProto[pb.EnrollmentLevel](mapCtx, in.EnrollmentLevel)
	return out
}
