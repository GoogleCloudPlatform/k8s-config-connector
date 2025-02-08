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

package oracledatabase

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/oracledatabase/apiv1/oracledatabasepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/oracledatabase/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func CloudAccountDetails_FromProto(mapCtx *direct.MapContext, in *pb.CloudAccountDetails) *krm.CloudAccountDetails {
	if in == nil {
		return nil
	}
	out := &krm.CloudAccountDetails{}
	// MISSING: CloudAccount
	// MISSING: CloudAccountHomeRegion
	// MISSING: LinkExistingAccountURI
	// MISSING: AccountCreationURI
	return out
}
func CloudAccountDetails_ToProto(mapCtx *direct.MapContext, in *krm.CloudAccountDetails) *pb.CloudAccountDetails {
	if in == nil {
		return nil
	}
	out := &pb.CloudAccountDetails{}
	// MISSING: CloudAccount
	// MISSING: CloudAccountHomeRegion
	// MISSING: LinkExistingAccountURI
	// MISSING: AccountCreationURI
	return out
}
func CloudAccountDetailsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CloudAccountDetails) *krm.CloudAccountDetailsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudAccountDetailsObservedState{}
	out.CloudAccount = direct.LazyPtr(in.GetCloudAccount())
	out.CloudAccountHomeRegion = direct.LazyPtr(in.GetCloudAccountHomeRegion())
	out.LinkExistingAccountURI = in.LinkExistingAccountUri
	out.AccountCreationURI = in.AccountCreationUri
	return out
}
func CloudAccountDetailsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudAccountDetailsObservedState) *pb.CloudAccountDetails {
	if in == nil {
		return nil
	}
	out := &pb.CloudAccountDetails{}
	out.CloudAccount = direct.ValueOf(in.CloudAccount)
	out.CloudAccountHomeRegion = direct.ValueOf(in.CloudAccountHomeRegion)
	out.LinkExistingAccountUri = in.LinkExistingAccountURI
	out.AccountCreationUri = in.AccountCreationURI
	return out
}
func Entitlement_FromProto(mapCtx *direct.MapContext, in *pb.Entitlement) *krm.Entitlement {
	if in == nil {
		return nil
	}
	out := &krm.Entitlement{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CloudAccountDetails = CloudAccountDetails_FromProto(mapCtx, in.GetCloudAccountDetails())
	// MISSING: EntitlementID
	// MISSING: State
	return out
}
func Entitlement_ToProto(mapCtx *direct.MapContext, in *krm.Entitlement) *pb.Entitlement {
	if in == nil {
		return nil
	}
	out := &pb.Entitlement{}
	out.Name = direct.ValueOf(in.Name)
	out.CloudAccountDetails = CloudAccountDetails_ToProto(mapCtx, in.CloudAccountDetails)
	// MISSING: EntitlementID
	// MISSING: State
	return out
}
func EntitlementObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Entitlement) *krm.EntitlementObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EntitlementObservedState{}
	// MISSING: Name
	out.CloudAccountDetails = CloudAccountDetailsObservedState_FromProto(mapCtx, in.GetCloudAccountDetails())
	out.EntitlementID = direct.LazyPtr(in.GetEntitlementId())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func EntitlementObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EntitlementObservedState) *pb.Entitlement {
	if in == nil {
		return nil
	}
	out := &pb.Entitlement{}
	// MISSING: Name
	out.CloudAccountDetails = CloudAccountDetailsObservedState_ToProto(mapCtx, in.CloudAccountDetails)
	out.EntitlementId = direct.ValueOf(in.EntitlementID)
	out.State = direct.Enum_ToProto[pb.Entitlement_State](mapCtx, in.State)
	return out
}
func OracledatabaseEntitlementObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Entitlement) *krm.OracledatabaseEntitlementObservedState {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseEntitlementObservedState{}
	// MISSING: Name
	// MISSING: CloudAccountDetails
	// MISSING: EntitlementID
	// MISSING: State
	return out
}
func OracledatabaseEntitlementObservedState_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseEntitlementObservedState) *pb.Entitlement {
	if in == nil {
		return nil
	}
	out := &pb.Entitlement{}
	// MISSING: Name
	// MISSING: CloudAccountDetails
	// MISSING: EntitlementID
	// MISSING: State
	return out
}
func OracledatabaseEntitlementSpec_FromProto(mapCtx *direct.MapContext, in *pb.Entitlement) *krm.OracledatabaseEntitlementSpec {
	if in == nil {
		return nil
	}
	out := &krm.OracledatabaseEntitlementSpec{}
	// MISSING: Name
	// MISSING: CloudAccountDetails
	// MISSING: EntitlementID
	// MISSING: State
	return out
}
func OracledatabaseEntitlementSpec_ToProto(mapCtx *direct.MapContext, in *krm.OracledatabaseEntitlementSpec) *pb.Entitlement {
	if in == nil {
		return nil
	}
	out := &pb.Entitlement{}
	// MISSING: Name
	// MISSING: CloudAccountDetails
	// MISSING: EntitlementID
	// MISSING: State
	return out
}
