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

package managedidentities

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/managedidentities/apiv1/managedidentitiespb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/managedidentities/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Domain_FromProto(mapCtx *direct.MapContext, in *pb.Domain) *krm.Domain {
	if in == nil {
		return nil
	}
	out := &krm.Domain{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Labels = in.Labels
	out.AuthorizedNetworks = in.AuthorizedNetworks
	out.ReservedIPRange = direct.LazyPtr(in.GetReservedIpRange())
	out.Locations = in.Locations
	out.Admin = direct.LazyPtr(in.GetAdmin())
	// MISSING: Fqdn
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StatusMessage
	// MISSING: Trusts
	return out
}
func Domain_ToProto(mapCtx *direct.MapContext, in *krm.Domain) *pb.Domain {
	if in == nil {
		return nil
	}
	out := &pb.Domain{}
	out.Name = direct.ValueOf(in.Name)
	out.Labels = in.Labels
	out.AuthorizedNetworks = in.AuthorizedNetworks
	out.ReservedIpRange = direct.ValueOf(in.ReservedIPRange)
	out.Locations = in.Locations
	out.Admin = direct.ValueOf(in.Admin)
	// MISSING: Fqdn
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StatusMessage
	// MISSING: Trusts
	return out
}
func DomainObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Domain) *krm.DomainObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DomainObservedState{}
	// MISSING: Name
	// MISSING: Labels
	// MISSING: AuthorizedNetworks
	// MISSING: ReservedIPRange
	// MISSING: Locations
	// MISSING: Admin
	out.Fqdn = direct.LazyPtr(in.GetFqdn())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StatusMessage = direct.LazyPtr(in.GetStatusMessage())
	out.Trusts = direct.Slice_FromProto(mapCtx, in.Trusts, Trust_FromProto)
	return out
}
func DomainObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DomainObservedState) *pb.Domain {
	if in == nil {
		return nil
	}
	out := &pb.Domain{}
	// MISSING: Name
	// MISSING: Labels
	// MISSING: AuthorizedNetworks
	// MISSING: ReservedIPRange
	// MISSING: Locations
	// MISSING: Admin
	out.Fqdn = direct.ValueOf(in.Fqdn)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Domain_State](mapCtx, in.State)
	out.StatusMessage = direct.ValueOf(in.StatusMessage)
	out.Trusts = direct.Slice_ToProto(mapCtx, in.Trusts, Trust_ToProto)
	return out
}
func ManagedidentitiesDomainObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Domain) *krm.ManagedidentitiesDomainObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ManagedidentitiesDomainObservedState{}
	// MISSING: Name
	// MISSING: Labels
	// MISSING: AuthorizedNetworks
	// MISSING: ReservedIPRange
	// MISSING: Locations
	// MISSING: Admin
	// MISSING: Fqdn
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StatusMessage
	// MISSING: Trusts
	return out
}
func ManagedidentitiesDomainObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ManagedidentitiesDomainObservedState) *pb.Domain {
	if in == nil {
		return nil
	}
	out := &pb.Domain{}
	// MISSING: Name
	// MISSING: Labels
	// MISSING: AuthorizedNetworks
	// MISSING: ReservedIPRange
	// MISSING: Locations
	// MISSING: Admin
	// MISSING: Fqdn
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StatusMessage
	// MISSING: Trusts
	return out
}
func ManagedidentitiesDomainSpec_FromProto(mapCtx *direct.MapContext, in *pb.Domain) *krm.ManagedidentitiesDomainSpec {
	if in == nil {
		return nil
	}
	out := &krm.ManagedidentitiesDomainSpec{}
	// MISSING: Name
	// MISSING: Labels
	// MISSING: AuthorizedNetworks
	// MISSING: ReservedIPRange
	// MISSING: Locations
	// MISSING: Admin
	// MISSING: Fqdn
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StatusMessage
	// MISSING: Trusts
	return out
}
func ManagedidentitiesDomainSpec_ToProto(mapCtx *direct.MapContext, in *krm.ManagedidentitiesDomainSpec) *pb.Domain {
	if in == nil {
		return nil
	}
	out := &pb.Domain{}
	// MISSING: Name
	// MISSING: Labels
	// MISSING: AuthorizedNetworks
	// MISSING: ReservedIPRange
	// MISSING: Locations
	// MISSING: Admin
	// MISSING: Fqdn
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StatusMessage
	// MISSING: Trusts
	return out
}
func Trust_FromProto(mapCtx *direct.MapContext, in *pb.Trust) *krm.Trust {
	if in == nil {
		return nil
	}
	out := &krm.Trust{}
	out.TargetDomainName = direct.LazyPtr(in.GetTargetDomainName())
	out.TrustType = direct.Enum_FromProto(mapCtx, in.GetTrustType())
	out.TrustDirection = direct.Enum_FromProto(mapCtx, in.GetTrustDirection())
	out.SelectiveAuthentication = direct.LazyPtr(in.GetSelectiveAuthentication())
	out.TargetDnsIPAddresses = in.TargetDnsIpAddresses
	out.TrustHandshakeSecret = direct.LazyPtr(in.GetTrustHandshakeSecret())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateDescription
	// MISSING: LastTrustHeartbeatTime
	return out
}
func Trust_ToProto(mapCtx *direct.MapContext, in *krm.Trust) *pb.Trust {
	if in == nil {
		return nil
	}
	out := &pb.Trust{}
	out.TargetDomainName = direct.ValueOf(in.TargetDomainName)
	out.TrustType = direct.Enum_ToProto[pb.Trust_TrustType](mapCtx, in.TrustType)
	out.TrustDirection = direct.Enum_ToProto[pb.Trust_TrustDirection](mapCtx, in.TrustDirection)
	out.SelectiveAuthentication = direct.ValueOf(in.SelectiveAuthentication)
	out.TargetDnsIpAddresses = in.TargetDnsIPAddresses
	out.TrustHandshakeSecret = direct.ValueOf(in.TrustHandshakeSecret)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: State
	// MISSING: StateDescription
	// MISSING: LastTrustHeartbeatTime
	return out
}
func TrustObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Trust) *krm.TrustObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TrustObservedState{}
	// MISSING: TargetDomainName
	// MISSING: TrustType
	// MISSING: TrustDirection
	// MISSING: SelectiveAuthentication
	// MISSING: TargetDnsIPAddresses
	// MISSING: TrustHandshakeSecret
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateDescription = direct.LazyPtr(in.GetStateDescription())
	out.LastTrustHeartbeatTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastTrustHeartbeatTime())
	return out
}
func TrustObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TrustObservedState) *pb.Trust {
	if in == nil {
		return nil
	}
	out := &pb.Trust{}
	// MISSING: TargetDomainName
	// MISSING: TrustType
	// MISSING: TrustDirection
	// MISSING: SelectiveAuthentication
	// MISSING: TargetDnsIPAddresses
	// MISSING: TrustHandshakeSecret
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.Trust_State](mapCtx, in.State)
	out.StateDescription = direct.ValueOf(in.StateDescription)
	out.LastTrustHeartbeatTime = direct.StringTimestamp_ToProto(mapCtx, in.LastTrustHeartbeatTime)
	return out
}
