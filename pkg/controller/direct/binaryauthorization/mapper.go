// Copyright 2026 Google LLC
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

// +generated:mapper
// krm.group: binaryauthorization.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.binaryauthorization.v1

package binaryauthorization

import (
	pb "cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/binaryauthorization/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AdmissionRule_RequireAttestationsBy_FromProto(mapCtx *direct.MapContext, in []string) []refsv1beta1.BinaryAuthorizationAttestorRef {
	if in == nil {
		return nil
	}
	out := make([]refsv1beta1.BinaryAuthorizationAttestorRef, len(in))
	for i, s := range in {
		out[i] = refsv1beta1.BinaryAuthorizationAttestorRef{
			External: s,
		}
	}
	return out
}

func AdmissionRule_RequireAttestationsBy_ToProto(mapCtx *direct.MapContext, in []refsv1beta1.BinaryAuthorizationAttestorRef) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	for i, ref := range in {
		if ref.External != "" {
			out[i] = ref.External
		} else {
			mapCtx.Errorf("RequireAttestationsBy[%d] has no external name (was it pre-resolved?)", i)
		}
	}
	return out
}

func BinaryAuthorizationPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krm.BinaryAuthorizationPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.BinaryAuthorizationPolicySpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.GlobalPolicyEvaluationMode = direct.Enum_FromProto(mapCtx, in.GetGlobalPolicyEvaluationMode())
	out.AdmissionWhitelistPatterns = direct.Slice_FromProto(mapCtx, in.AdmissionWhitelistPatterns, AdmissionWhitelistPattern_FromProto)
	out.ClusterAdmissionRules = BinaryAuthorizationPolicySpec_ClusterAdmissionRules_FromProto(mapCtx, in.GetClusterAdmissionRules())
	out.KubernetesNamespaceAdmissionRules = BinaryAuthorizationPolicySpec_KubernetesNamespaceAdmissionRules_FromProto(mapCtx, in.GetKubernetesNamespaceAdmissionRules())
	out.KubernetesServiceAccountAdmissionRules = BinaryAuthorizationPolicySpec_KubernetesServiceAccountAdmissionRules_FromProto(mapCtx, in.GetKubernetesServiceAccountAdmissionRules())
	out.IstioServiceIdentityAdmissionRules = BinaryAuthorizationPolicySpec_IstioServiceIdentityAdmissionRules_FromProto(mapCtx, in.GetIstioServiceIdentityAdmissionRules())
	out.DefaultAdmissionRule = AdmissionRule_FromProto(mapCtx, in.GetDefaultAdmissionRule())
	return out
}

func BinaryAuthorizationPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.BinaryAuthorizationPolicySpec) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	out.Description = direct.ValueOf(in.Description)
	out.GlobalPolicyEvaluationMode = direct.Enum_ToProto[pb.Policy_GlobalPolicyEvaluationMode](mapCtx, in.GlobalPolicyEvaluationMode)
	out.AdmissionWhitelistPatterns = direct.Slice_ToProto(mapCtx, in.AdmissionWhitelistPatterns, AdmissionWhitelistPattern_ToProto)
	out.ClusterAdmissionRules = BinaryAuthorizationPolicySpec_ClusterAdmissionRules_ToProto(mapCtx, in.ClusterAdmissionRules)
	out.KubernetesNamespaceAdmissionRules = BinaryAuthorizationPolicySpec_KubernetesNamespaceAdmissionRules_ToProto(mapCtx, in.KubernetesNamespaceAdmissionRules)
	out.KubernetesServiceAccountAdmissionRules = BinaryAuthorizationPolicySpec_KubernetesServiceAccountAdmissionRules_ToProto(mapCtx, in.KubernetesServiceAccountAdmissionRules)
	out.IstioServiceIdentityAdmissionRules = BinaryAuthorizationPolicySpec_IstioServiceIdentityAdmissionRules_ToProto(mapCtx, in.IstioServiceIdentityAdmissionRules)
	out.DefaultAdmissionRule = AdmissionRule_ToProto(mapCtx, in.DefaultAdmissionRule)
	return out
}

func BinaryAuthorizationPolicySpec_ClusterAdmissionRules_FromProto(mapCtx *direct.MapContext, in map[string]*pb.AdmissionRule) map[string]krm.AdmissionRule {
	if in == nil {
		return nil
	}
	out := make(map[string]krm.AdmissionRule)
	for k, v := range in {
		out[k] = *AdmissionRule_FromProto(mapCtx, v)
	}
	return out
}

func BinaryAuthorizationPolicySpec_ClusterAdmissionRules_ToProto(mapCtx *direct.MapContext, in map[string]krm.AdmissionRule) map[string]*pb.AdmissionRule {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.AdmissionRule)
	for k, v := range in {
		v := v
		out[k] = AdmissionRule_ToProto(mapCtx, &v)
	}
	return out
}

func BinaryAuthorizationPolicySpec_KubernetesNamespaceAdmissionRules_FromProto(mapCtx *direct.MapContext, in map[string]*pb.AdmissionRule) map[string]krm.AdmissionRule {
	if in == nil {
		return nil
	}
	out := make(map[string]krm.AdmissionRule)
	for k, v := range in {
		out[k] = *AdmissionRule_FromProto(mapCtx, v)
	}
	return out
}

func BinaryAuthorizationPolicySpec_KubernetesNamespaceAdmissionRules_ToProto(mapCtx *direct.MapContext, in map[string]krm.AdmissionRule) map[string]*pb.AdmissionRule {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.AdmissionRule)
	for k, v := range in {
		v := v
		out[k] = AdmissionRule_ToProto(mapCtx, &v)
	}
	return out
}

func BinaryAuthorizationPolicySpec_KubernetesServiceAccountAdmissionRules_FromProto(mapCtx *direct.MapContext, in map[string]*pb.AdmissionRule) map[string]krm.AdmissionRule {
	if in == nil {
		return nil
	}
	out := make(map[string]krm.AdmissionRule)
	for k, v := range in {
		out[k] = *AdmissionRule_FromProto(mapCtx, v)
	}
	return out
}

func BinaryAuthorizationPolicySpec_KubernetesServiceAccountAdmissionRules_ToProto(mapCtx *direct.MapContext, in map[string]krm.AdmissionRule) map[string]*pb.AdmissionRule {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.AdmissionRule)
	for k, v := range in {
		v := v
		out[k] = AdmissionRule_ToProto(mapCtx, &v)
	}
	return out
}

func BinaryAuthorizationPolicySpec_IstioServiceIdentityAdmissionRules_FromProto(mapCtx *direct.MapContext, in map[string]*pb.AdmissionRule) map[string]krm.AdmissionRule {
	if in == nil {
		return nil
	}
	out := make(map[string]krm.AdmissionRule)
	for k, v := range in {
		out[k] = *AdmissionRule_FromProto(mapCtx, v)
	}
	return out
}

func BinaryAuthorizationPolicySpec_IstioServiceIdentityAdmissionRules_ToProto(mapCtx *direct.MapContext, in map[string]krm.AdmissionRule) map[string]*pb.AdmissionRule {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.AdmissionRule)
	for k, v := range in {
		v := v
		out[k] = AdmissionRule_ToProto(mapCtx, &v)
	}
	return out
}

func BinaryAuthorizationPolicyStatus_FromProto(mapCtx *direct.MapContext, in *pb.Policy) *krm.BinaryAuthorizationPolicyStatus {
	if in == nil {
		return nil
	}
	out := &krm.BinaryAuthorizationPolicyStatus{}
	out.SelfLink = direct.LazyPtr(in.GetName())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}

func BinaryAuthorizationPolicyStatus_ToProto(mapCtx *direct.MapContext, in *krm.BinaryAuthorizationPolicyStatus) *pb.Policy {
	if in == nil {
		return nil
	}
	out := &pb.Policy{}
	out.Name = direct.ValueOf(in.SelfLink)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
