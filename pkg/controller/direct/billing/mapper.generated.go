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

package billing

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/billing/apiv1/billingpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/billing/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BillingBillingAccountObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BillingAccount) *krm.BillingBillingAccountObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BillingBillingAccountObservedState{}
	// MISSING: Name
	// MISSING: Open
	// MISSING: DisplayName
	// MISSING: MasterBillingAccount
	// MISSING: Parent
	// MISSING: CurrencyCode
	return out
}
func BillingBillingAccountObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BillingBillingAccountObservedState) *pb.BillingAccount {
	if in == nil {
		return nil
	}
	out := &pb.BillingAccount{}
	// MISSING: Name
	// MISSING: Open
	// MISSING: DisplayName
	// MISSING: MasterBillingAccount
	// MISSING: Parent
	// MISSING: CurrencyCode
	return out
}
func BillingBillingAccountSpec_FromProto(mapCtx *direct.MapContext, in *pb.BillingAccount) *krm.BillingBillingAccountSpec {
	if in == nil {
		return nil
	}
	out := &krm.BillingBillingAccountSpec{}
	// MISSING: Name
	// MISSING: Open
	// MISSING: DisplayName
	// MISSING: MasterBillingAccount
	// MISSING: Parent
	// MISSING: CurrencyCode
	return out
}
func BillingBillingAccountSpec_ToProto(mapCtx *direct.MapContext, in *krm.BillingBillingAccountSpec) *pb.BillingAccount {
	if in == nil {
		return nil
	}
	out := &pb.BillingAccount{}
	// MISSING: Name
	// MISSING: Open
	// MISSING: DisplayName
	// MISSING: MasterBillingAccount
	// MISSING: Parent
	// MISSING: CurrencyCode
	return out
}
func BillingProjectBillingInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ProjectBillingInfo) *krm.BillingProjectBillingInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BillingProjectBillingInfoObservedState{}
	// MISSING: Name
	// MISSING: ProjectID
	// MISSING: BillingAccountName
	// MISSING: BillingEnabled
	return out
}
func BillingProjectBillingInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BillingProjectBillingInfoObservedState) *pb.ProjectBillingInfo {
	if in == nil {
		return nil
	}
	out := &pb.ProjectBillingInfo{}
	// MISSING: Name
	// MISSING: ProjectID
	// MISSING: BillingAccountName
	// MISSING: BillingEnabled
	return out
}
func BillingProjectBillingInfoSpec_FromProto(mapCtx *direct.MapContext, in *pb.ProjectBillingInfo) *krm.BillingProjectBillingInfoSpec {
	if in == nil {
		return nil
	}
	out := &krm.BillingProjectBillingInfoSpec{}
	// MISSING: Name
	// MISSING: ProjectID
	// MISSING: BillingAccountName
	// MISSING: BillingEnabled
	return out
}
func BillingProjectBillingInfoSpec_ToProto(mapCtx *direct.MapContext, in *krm.BillingProjectBillingInfoSpec) *pb.ProjectBillingInfo {
	if in == nil {
		return nil
	}
	out := &pb.ProjectBillingInfo{}
	// MISSING: Name
	// MISSING: ProjectID
	// MISSING: BillingAccountName
	// MISSING: BillingEnabled
	return out
}
func ProjectBillingInfo_FromProto(mapCtx *direct.MapContext, in *pb.ProjectBillingInfo) *krm.ProjectBillingInfo {
	if in == nil {
		return nil
	}
	out := &krm.ProjectBillingInfo{}
	// MISSING: Name
	// MISSING: ProjectID
	out.BillingAccountName = direct.LazyPtr(in.GetBillingAccountName())
	// MISSING: BillingEnabled
	return out
}
func ProjectBillingInfo_ToProto(mapCtx *direct.MapContext, in *krm.ProjectBillingInfo) *pb.ProjectBillingInfo {
	if in == nil {
		return nil
	}
	out := &pb.ProjectBillingInfo{}
	// MISSING: Name
	// MISSING: ProjectID
	out.BillingAccountName = direct.ValueOf(in.BillingAccountName)
	// MISSING: BillingEnabled
	return out
}
func ProjectBillingInfoObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ProjectBillingInfo) *krm.ProjectBillingInfoObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ProjectBillingInfoObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	// MISSING: BillingAccountName
	out.BillingEnabled = direct.LazyPtr(in.GetBillingEnabled())
	return out
}
func ProjectBillingInfoObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ProjectBillingInfoObservedState) *pb.ProjectBillingInfo {
	if in == nil {
		return nil
	}
	out := &pb.ProjectBillingInfo{}
	out.Name = direct.ValueOf(in.Name)
	out.ProjectId = direct.ValueOf(in.ProjectID)
	// MISSING: BillingAccountName
	out.BillingEnabled = direct.ValueOf(in.BillingEnabled)
	return out
}
