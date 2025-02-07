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
func AccessApprovalServiceAccount_FromProto(mapCtx *direct.MapContext, in *pb.AccessApprovalServiceAccount) *krm.AccessApprovalServiceAccount {
	if in == nil {
		return nil
	}
	out := &krm.AccessApprovalServiceAccount{}
	out.Name = direct.LazyPtr(in.GetName())
	out.AccountEmail = direct.LazyPtr(in.GetAccountEmail())
	return out
}
func AccessApprovalServiceAccount_ToProto(mapCtx *direct.MapContext, in *krm.AccessApprovalServiceAccount) *pb.AccessApprovalServiceAccount {
	if in == nil {
		return nil
	}
	out := &pb.AccessApprovalServiceAccount{}
	out.Name = direct.ValueOf(in.Name)
	out.AccountEmail = direct.ValueOf(in.AccountEmail)
	return out
}
func AccessapprovalAccessApprovalServiceAccountObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AccessApprovalServiceAccount) *krm.AccessapprovalAccessApprovalServiceAccountObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AccessapprovalAccessApprovalServiceAccountObservedState{}
	// MISSING: Name
	// MISSING: AccountEmail
	return out
}
func AccessapprovalAccessApprovalServiceAccountObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AccessapprovalAccessApprovalServiceAccountObservedState) *pb.AccessApprovalServiceAccount {
	if in == nil {
		return nil
	}
	out := &pb.AccessApprovalServiceAccount{}
	// MISSING: Name
	// MISSING: AccountEmail
	return out
}
func AccessapprovalAccessApprovalServiceAccountSpec_FromProto(mapCtx *direct.MapContext, in *pb.AccessApprovalServiceAccount) *krm.AccessapprovalAccessApprovalServiceAccountSpec {
	if in == nil {
		return nil
	}
	out := &krm.AccessapprovalAccessApprovalServiceAccountSpec{}
	// MISSING: Name
	// MISSING: AccountEmail
	return out
}
func AccessapprovalAccessApprovalServiceAccountSpec_ToProto(mapCtx *direct.MapContext, in *krm.AccessapprovalAccessApprovalServiceAccountSpec) *pb.AccessApprovalServiceAccount {
	if in == nil {
		return nil
	}
	out := &pb.AccessApprovalServiceAccount{}
	// MISSING: Name
	// MISSING: AccountEmail
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
