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

package iam

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/iam/apiv1beta/iampb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1alpha1"
)
func IamWorkloadIdentityPoolProviderObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadIdentityPoolProvider) *krm.IamWorkloadIdentityPoolProviderObservedState {
	if in == nil {
		return nil
	}
	out := &krm.IamWorkloadIdentityPoolProviderObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: State
	// MISSING: Disabled
	// MISSING: AttributeMapping
	// MISSING: AttributeCondition
	// MISSING: Aws
	// MISSING: Oidc
	return out
}
func IamWorkloadIdentityPoolProviderObservedState_ToProto(mapCtx *direct.MapContext, in *krm.IamWorkloadIdentityPoolProviderObservedState) *pb.WorkloadIdentityPoolProvider {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadIdentityPoolProvider{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: State
	// MISSING: Disabled
	// MISSING: AttributeMapping
	// MISSING: AttributeCondition
	// MISSING: Aws
	// MISSING: Oidc
	return out
}
func IamWorkloadIdentityPoolProviderSpec_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadIdentityPoolProvider) *krm.IamWorkloadIdentityPoolProviderSpec {
	if in == nil {
		return nil
	}
	out := &krm.IamWorkloadIdentityPoolProviderSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: State
	// MISSING: Disabled
	// MISSING: AttributeMapping
	// MISSING: AttributeCondition
	// MISSING: Aws
	// MISSING: Oidc
	return out
}
func IamWorkloadIdentityPoolProviderSpec_ToProto(mapCtx *direct.MapContext, in *krm.IamWorkloadIdentityPoolProviderSpec) *pb.WorkloadIdentityPoolProvider {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadIdentityPoolProvider{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: State
	// MISSING: Disabled
	// MISSING: AttributeMapping
	// MISSING: AttributeCondition
	// MISSING: Aws
	// MISSING: Oidc
	return out
}
func WorkloadIdentityPoolProvider_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadIdentityPoolProvider) *krm.WorkloadIdentityPoolProvider {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadIdentityPoolProvider{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: State
	out.Disabled = direct.LazyPtr(in.GetDisabled())
	out.AttributeMapping = in.AttributeMapping
	out.AttributeCondition = direct.LazyPtr(in.GetAttributeCondition())
	out.Aws = WorkloadIdentityPoolProvider_Aws_FromProto(mapCtx, in.GetAws())
	out.Oidc = WorkloadIdentityPoolProvider_Oidc_FromProto(mapCtx, in.GetOidc())
	return out
}
func WorkloadIdentityPoolProvider_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadIdentityPoolProvider) *pb.WorkloadIdentityPoolProvider {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadIdentityPoolProvider{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: State
	out.Disabled = direct.ValueOf(in.Disabled)
	out.AttributeMapping = in.AttributeMapping
	out.AttributeCondition = direct.ValueOf(in.AttributeCondition)
	if oneof := WorkloadIdentityPoolProvider_Aws_ToProto(mapCtx, in.Aws); oneof != nil {
		out.ProviderConfig = &pb.WorkloadIdentityPoolProvider_Aws_{Aws: oneof}
	}
	if oneof := WorkloadIdentityPoolProvider_Oidc_ToProto(mapCtx, in.Oidc); oneof != nil {
		out.ProviderConfig = &pb.WorkloadIdentityPoolProvider_Oidc_{Oidc: oneof}
	}
	return out
}
func WorkloadIdentityPoolProviderObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadIdentityPoolProvider) *krm.WorkloadIdentityPoolProviderObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadIdentityPoolProviderObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Disabled
	// MISSING: AttributeMapping
	// MISSING: AttributeCondition
	// MISSING: Aws
	// MISSING: Oidc
	return out
}
func WorkloadIdentityPoolProviderObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadIdentityPoolProviderObservedState) *pb.WorkloadIdentityPoolProvider {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadIdentityPoolProvider{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Description
	out.State = direct.Enum_ToProto[pb.WorkloadIdentityPoolProvider_State](mapCtx, in.State)
	// MISSING: Disabled
	// MISSING: AttributeMapping
	// MISSING: AttributeCondition
	// MISSING: Aws
	// MISSING: Oidc
	return out
}
func WorkloadIdentityPoolProvider_Aws_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadIdentityPoolProvider_Aws) *krm.WorkloadIdentityPoolProvider_Aws {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadIdentityPoolProvider_Aws{}
	out.AccountID = direct.LazyPtr(in.GetAccountId())
	return out
}
func WorkloadIdentityPoolProvider_Aws_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadIdentityPoolProvider_Aws) *pb.WorkloadIdentityPoolProvider_Aws {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadIdentityPoolProvider_Aws{}
	out.AccountId = direct.ValueOf(in.AccountID)
	return out
}
func WorkloadIdentityPoolProvider_Oidc_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadIdentityPoolProvider_Oidc) *krm.WorkloadIdentityPoolProvider_Oidc {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadIdentityPoolProvider_Oidc{}
	out.IssuerURI = direct.LazyPtr(in.GetIssuerUri())
	out.AllowedAudiences = in.AllowedAudiences
	return out
}
func WorkloadIdentityPoolProvider_Oidc_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadIdentityPoolProvider_Oidc) *pb.WorkloadIdentityPoolProvider_Oidc {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadIdentityPoolProvider_Oidc{}
	out.IssuerUri = direct.ValueOf(in.IssuerURI)
	out.AllowedAudiences = in.AllowedAudiences
	return out
}
