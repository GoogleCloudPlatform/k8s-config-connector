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

package recaptchaenterprise

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/recaptchaenterprise/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func RecaptchaenterpriseRelatedAccountGroupMembershipObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RelatedAccountGroupMembership) *krm.RecaptchaenterpriseRelatedAccountGroupMembershipObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RecaptchaenterpriseRelatedAccountGroupMembershipObservedState{}
	// MISSING: Name
	// MISSING: AccountID
	// MISSING: HashedAccountID
	return out
}
func RecaptchaenterpriseRelatedAccountGroupMembershipObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RecaptchaenterpriseRelatedAccountGroupMembershipObservedState) *pb.RelatedAccountGroupMembership {
	if in == nil {
		return nil
	}
	out := &pb.RelatedAccountGroupMembership{}
	// MISSING: Name
	// MISSING: AccountID
	// MISSING: HashedAccountID
	return out
}
func RecaptchaenterpriseRelatedAccountGroupMembershipSpec_FromProto(mapCtx *direct.MapContext, in *pb.RelatedAccountGroupMembership) *krm.RecaptchaenterpriseRelatedAccountGroupMembershipSpec {
	if in == nil {
		return nil
	}
	out := &krm.RecaptchaenterpriseRelatedAccountGroupMembershipSpec{}
	// MISSING: Name
	// MISSING: AccountID
	// MISSING: HashedAccountID
	return out
}
func RecaptchaenterpriseRelatedAccountGroupMembershipSpec_ToProto(mapCtx *direct.MapContext, in *krm.RecaptchaenterpriseRelatedAccountGroupMembershipSpec) *pb.RelatedAccountGroupMembership {
	if in == nil {
		return nil
	}
	out := &pb.RelatedAccountGroupMembership{}
	// MISSING: Name
	// MISSING: AccountID
	// MISSING: HashedAccountID
	return out
}
func RelatedAccountGroupMembership_FromProto(mapCtx *direct.MapContext, in *pb.RelatedAccountGroupMembership) *krm.RelatedAccountGroupMembership {
	if in == nil {
		return nil
	}
	out := &krm.RelatedAccountGroupMembership{}
	out.Name = direct.LazyPtr(in.GetName())
	out.AccountID = direct.LazyPtr(in.GetAccountId())
	out.HashedAccountID = in.GetHashedAccountId()
	return out
}
func RelatedAccountGroupMembership_ToProto(mapCtx *direct.MapContext, in *krm.RelatedAccountGroupMembership) *pb.RelatedAccountGroupMembership {
	if in == nil {
		return nil
	}
	out := &pb.RelatedAccountGroupMembership{}
	out.Name = direct.ValueOf(in.Name)
	out.AccountId = direct.ValueOf(in.AccountID)
	out.HashedAccountId = in.HashedAccountID
	return out
}
