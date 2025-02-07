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

package channel

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/channel/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/channel/apiv1/channelpb"
)
func BillingAccount_FromProto(mapCtx *direct.MapContext, in *pb.BillingAccount) *krm.BillingAccount {
	if in == nil {
		return nil
	}
	out := &krm.BillingAccount{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: CreateTime
	// MISSING: CurrencyCode
	// MISSING: RegionCode
	return out
}
func BillingAccount_ToProto(mapCtx *direct.MapContext, in *krm.BillingAccount) *pb.BillingAccount {
	if in == nil {
		return nil
	}
	out := &pb.BillingAccount{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: CreateTime
	// MISSING: CurrencyCode
	// MISSING: RegionCode
	return out
}
func BillingAccountObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BillingAccount) *krm.BillingAccountObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BillingAccountObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.CurrencyCode = direct.LazyPtr(in.GetCurrencyCode())
	out.RegionCode = direct.LazyPtr(in.GetRegionCode())
	return out
}
func BillingAccountObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BillingAccountObservedState) *pb.BillingAccount {
	if in == nil {
		return nil
	}
	out := &pb.BillingAccount{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.CurrencyCode = direct.ValueOf(in.CurrencyCode)
	out.RegionCode = direct.ValueOf(in.RegionCode)
	return out
}
func ChannelBillingAccountObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BillingAccount) *krm.ChannelBillingAccountObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ChannelBillingAccountObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: CurrencyCode
	// MISSING: RegionCode
	return out
}
func ChannelBillingAccountObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ChannelBillingAccountObservedState) *pb.BillingAccount {
	if in == nil {
		return nil
	}
	out := &pb.BillingAccount{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: CurrencyCode
	// MISSING: RegionCode
	return out
}
func ChannelBillingAccountSpec_FromProto(mapCtx *direct.MapContext, in *pb.BillingAccount) *krm.ChannelBillingAccountSpec {
	if in == nil {
		return nil
	}
	out := &krm.ChannelBillingAccountSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: CurrencyCode
	// MISSING: RegionCode
	return out
}
func ChannelBillingAccountSpec_ToProto(mapCtx *direct.MapContext, in *krm.ChannelBillingAccountSpec) *pb.BillingAccount {
	if in == nil {
		return nil
	}
	out := &pb.BillingAccount{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: CreateTime
	// MISSING: CurrencyCode
	// MISSING: RegionCode
	return out
}
