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

package binaryauthorization

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/binaryauthorization/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/binaryauthorization/apiv1beta1/binaryauthorizationpb"
)
func AdmissionRule_FromProto(mapCtx *direct.MapContext, in *pb.AdmissionRule) *krm.AdmissionRule {
	if in == nil {
		return nil
	}
	out := &krm.AdmissionRule{}
	out.EvaluationMode = direct.Enum_FromProto(mapCtx, in.GetEvaluationMode())
	out.RequireAttestationsBy = in.RequireAttestationsBy
	out.EnforcementMode = direct.Enum_FromProto(mapCtx, in.GetEnforcementMode())
	return out
}
func AdmissionRule_ToProto(mapCtx *direct.MapContext, in *krm.AdmissionRule) *pb.AdmissionRule {
	if in == nil {
		return nil
	}
	out := &pb.AdmissionRule{}
	out.EvaluationMode = direct.Enum_ToProto[pb.AdmissionRule_EvaluationMode](mapCtx, in.EvaluationMode)
	out.RequireAttestationsBy = in.RequireAttestationsBy
	out.EnforcementMode = direct.Enum_ToProto[pb.AdmissionRule_EnforcementMode](mapCtx, in.EnforcementMode)
	return out
}
func AdmissionWhitelistPattern_FromProto(mapCtx *direct.MapContext, in *pb.AdmissionWhitelistPattern) *krm.AdmissionWhitelistPattern {
	if in == nil {
		return nil
	}
	out := &krm.AdmissionWhitelistPattern{}
	out.NamePattern = direct.LazyPtr(in.GetNamePattern())
	return out
}
func AdmissionWhitelistPattern_ToProto(mapCtx *direct.MapContext, in *krm.AdmissionWhitelistPattern) *pb.AdmissionWhitelistPattern {
	if in == nil {
		return nil
	}
	out := &pb.AdmissionWhitelistPattern{}
	out.NamePattern = direct.ValueOf(in.NamePattern)
	return out
}
func Policy_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krm.Policy {
	if in == nil {
		return nil
	}
	out := &krm.Policy{}
	// MISSING: Name
	out.Description = direct.LazyPtr(in.GetDescription())
	out.GlobalPolicyEvaluationMode = direct.Enum_FromProto(mapCtx, in.GetGlobalPolicyEvaluationMode())
	out.AdmissionWhitelistPatterns = direct.Slice_FromProto(mapCtx, in.AdmissionWhitelistPatterns, AdmissionWhitelistPattern_FromProto)
	// MISSING: ClusterAdmissionRules
	// MISSING: KubernetesNamespaceAdmissionRules
	// MISSING: KubernetesServiceAccountAdmissionRules
	// MISSING: IstioServiceIdentityAdmissionRules
	out.DefaultAdmissionRule = AdmissionRule_FromProto(mapCtx, in.GetDefaultAdmissionRule())
	// MISSING: UpdateTime
	return out
}
func Policy_ToProto(mapCtx *direct.MapContext, in *krm.Policy) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	// MISSING: Name
	out.Description = direct.ValueOf(in.Description)
	out.GlobalPolicyEvaluationMode = direct.Enum_ToProto[pb.Policy_GlobalPolicyEvaluationMode](mapCtx, in.GlobalPolicyEvaluationMode)
	out.AdmissionWhitelistPatterns = direct.Slice_ToProto(mapCtx, in.AdmissionWhitelistPatterns, AdmissionWhitelistPattern_ToProto)
	// MISSING: ClusterAdmissionRules
	// MISSING: KubernetesNamespaceAdmissionRules
	// MISSING: KubernetesServiceAccountAdmissionRules
	// MISSING: IstioServiceIdentityAdmissionRules
	out.DefaultAdmissionRule = AdmissionRule_ToProto(mapCtx, in.DefaultAdmissionRule)
	// MISSING: UpdateTime
	return out
}
func PolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krm.PolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PolicyObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Description
	// MISSING: GlobalPolicyEvaluationMode
	// MISSING: AdmissionWhitelistPatterns
	// MISSING: ClusterAdmissionRules
	// MISSING: KubernetesNamespaceAdmissionRules
	// MISSING: KubernetesServiceAccountAdmissionRules
	// MISSING: IstioServiceIdentityAdmissionRules
	// MISSING: DefaultAdmissionRule
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func PolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PolicyObservedState) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Description
	// MISSING: GlobalPolicyEvaluationMode
	// MISSING: AdmissionWhitelistPatterns
	// MISSING: ClusterAdmissionRules
	// MISSING: KubernetesNamespaceAdmissionRules
	// MISSING: KubernetesServiceAccountAdmissionRules
	// MISSING: IstioServiceIdentityAdmissionRules
	// MISSING: DefaultAdmissionRule
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
