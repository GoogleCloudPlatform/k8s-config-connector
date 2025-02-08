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
func RecaptchaenterpriseRelatedAccountGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.RelatedAccountGroup) *krm.RecaptchaenterpriseRelatedAccountGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RecaptchaenterpriseRelatedAccountGroupObservedState{}
	// MISSING: Name
	return out
}
func RecaptchaenterpriseRelatedAccountGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RecaptchaenterpriseRelatedAccountGroupObservedState) *pb.RelatedAccountGroup {
	if in == nil {
		return nil
	}
	out := &pb.RelatedAccountGroup{}
	// MISSING: Name
	return out
}
func RecaptchaenterpriseRelatedAccountGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.RelatedAccountGroup) *krm.RecaptchaenterpriseRelatedAccountGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.RecaptchaenterpriseRelatedAccountGroupSpec{}
	// MISSING: Name
	return out
}
func RecaptchaenterpriseRelatedAccountGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.RecaptchaenterpriseRelatedAccountGroupSpec) *pb.RelatedAccountGroup {
	if in == nil {
		return nil
	}
	out := &pb.RelatedAccountGroup{}
	// MISSING: Name
	return out
}
func RelatedAccountGroup_FromProto(mapCtx *direct.MapContext, in *pb.RelatedAccountGroup) *krm.RelatedAccountGroup {
	if in == nil {
		return nil
	}
	out := &krm.RelatedAccountGroup{}
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func RelatedAccountGroup_ToProto(mapCtx *direct.MapContext, in *krm.RelatedAccountGroup) *pb.RelatedAccountGroup {
	if in == nil {
		return nil
	}
	out := &pb.RelatedAccountGroup{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
