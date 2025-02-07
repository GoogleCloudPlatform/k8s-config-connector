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
func BillingServiceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krm.BillingServiceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BillingServiceObservedState{}
	// MISSING: Name
	// MISSING: ServiceID
	// MISSING: DisplayName
	// MISSING: BusinessEntityName
	return out
}
func BillingServiceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BillingServiceObservedState) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	// MISSING: Name
	// MISSING: ServiceID
	// MISSING: DisplayName
	// MISSING: BusinessEntityName
	return out
}
func BillingServiceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krm.BillingServiceSpec {
	if in == nil {
		return nil
	}
	out := &krm.BillingServiceSpec{}
	// MISSING: Name
	// MISSING: ServiceID
	// MISSING: DisplayName
	// MISSING: BusinessEntityName
	return out
}
func BillingServiceSpec_ToProto(mapCtx *direct.MapContext, in *krm.BillingServiceSpec) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	// MISSING: Name
	// MISSING: ServiceID
	// MISSING: DisplayName
	// MISSING: BusinessEntityName
	return out
}
func Service_FromProto(mapCtx *direct.MapContext, in *pb.Service) *krm.Service {
	if in == nil {
		return nil
	}
	out := &krm.Service{}
	out.Name = direct.LazyPtr(in.GetName())
	out.ServiceID = direct.LazyPtr(in.GetServiceId())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.BusinessEntityName = direct.LazyPtr(in.GetBusinessEntityName())
	return out
}
func Service_ToProto(mapCtx *direct.MapContext, in *krm.Service) *pb.Service {
	if in == nil {
		return nil
	}
	out := &pb.Service{}
	out.Name = direct.ValueOf(in.Name)
	out.ServiceId = direct.ValueOf(in.ServiceID)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.BusinessEntityName = direct.ValueOf(in.BusinessEntityName)
	return out
}
